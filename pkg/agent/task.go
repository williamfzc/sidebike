package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/williamfzc/sidebike/pkg/server"
)

func (agent *Agent) GetUrlTaskRequest() (*url.URL, error) {
	ret, err := url.Parse(fmt.Sprintf("http://%s%s", agent.GetRegistry(), server.AssignTask.Path))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (agent *Agent) GetUrlTaskDone() (*url.URL, error) {
	ret, err := url.Parse(fmt.Sprintf("http://%s%s", agent.GetRegistry(), server.DoneTask.Path))
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (agent *Agent) StartTaskRequestMonitor() {
	for {
		<-agent.taskRequestQueue
		agent.triggerTaskRequest()
	}
}

func (agent *Agent) triggerTaskRequest() {
	finalUrl, err := agent.GetUrlTaskRequest()
	if err != nil {
		logger.Errorf("failed to gen task url: %s", err)
		return
	}

	// todo: offer enough info for server to make decision
	requestJson := &server.TaskAssignRequest{
		MachineLabel: agent.MachineLabel,
	}
	jsonStr, _ := json.Marshal(requestJson)

	response, err := http.Post(finalUrl.String(), "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		logger.Errorf("request error: %s", err)
		return
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		logger.Errorf("read body error: %s", err)
		return
	}

	responseObj := &server.TaskResponse{}
	err = json.Unmarshal(data, responseObj)
	if err != nil {
		logger.Errorf("invalid response format: %s", err)
		return
	}
	if responseObj.Signal != server.SignalOk {
		logger.Errorf("status not ok")
		return
	}

	task := responseObj.Data
	logger.Infof("ready to run task %v", task)
	agent.taskTodoQueue <- &task
}
