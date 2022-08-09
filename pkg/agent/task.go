package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/williamfzc/sidebike/pkg/server"
	"io"
	"net/http"
	"net/url"
)

func (agent *Agent) GetTaskRequestUrl() (*url.URL, error) {
	ret, err := url.Parse(fmt.Sprintf("http://%s%s", agent.GetRegistry(), server.AssignTask.Path))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (agent *Agent) taskRequestMonitor() {
	for {
		<-agent.taskRequestQueue

		finalUrl, err := agent.GetTaskRequestUrl()
		if err != nil {
			logger.Errorf("failed to gen task url: %s", err)
			continue
		}

		// todo: offer enough info for server to make decision
		requestJson := &server.TaskAssign{
			MachinePath: agent.GetMachinePath(),
		}
		jsonStr, _ := json.Marshal(requestJson)

		response, err := http.Post(finalUrl.String(), "application/json", bytes.NewBuffer(jsonStr))
		if err != nil {
			logger.Errorf("request error: %s", err)
			continue
		}

		data, err := io.ReadAll(response.Body)
		if err != nil {
			logger.Errorf("read body error: %s", err)
			continue
		}

		responseObj := &server.TaskResponse{}
		err = json.Unmarshal(data, responseObj)
		if err != nil {
			logger.Errorf("invalid response format: %s", err)
			continue
		}
		if responseObj.Signal != server.SignalOk {
			logger.Errorf("status not ok")
			continue
		}

		task := responseObj.Data
		logger.Infof("ready to run task %s", task)
		agent.taskTodoQueue <- &task
	}
}
