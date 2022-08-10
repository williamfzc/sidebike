package agent

import (
	"bytes"
	"encoding/json"
	"github.com/go-cmd/cmd"
	"github.com/williamfzc/sidebike/pkg/server"
	"io"
	"net/http"
	"time"
)

const ResultLineLimit = 30

func (agent *Agent) taskWorkMonitor() {
	for {
		task := <-agent.taskTodoQueue

		commandToRun := task.Detail.Command
		ok := verifyCommand(commandToRun)
		if !ok {
			task.Status = server.TaskStatusError
			task.Detail.Result = []string{"command verify error"}
			agent.UpdateTaskStatus(task)
			continue
		}

		logger.Infof("run shell: %s", commandToRun)
		userCmd := cmd.NewCmd("bash", "-c", commandToRun)

		workspace := task.Detail.Workspace
		if workspace != "" {
			userCmd.Dir = workspace
		}

		go func() {
			<-time.After(time.Duration(task.Detail.Timeout) * time.Second)
			_ = userCmd.Stop()
		}()
		status := <-userCmd.Start()
		logger.Infof("task done: %v", status)

		if status.Exit == 0 {
			task.Status = server.TaskStatusFinished
		} else {
			task.Status = server.TaskStatusError
		}

		stdoutRet := userCmd.Status().Stdout
		if len(stdoutRet) > ResultLineLimit {
			stdoutRet = stdoutRet[ResultLineLimit:]
		}
		task.Detail.Result = stdoutRet
		agent.UpdateTaskStatus(task)
	}
}

func (agent *Agent) UpdateTaskStatus(task *server.Task) {
	// UpdateTaskStatus
	finalUrl, err := agent.GetUrlTaskDone()
	if err != nil {
		logger.Errorf("failed to gen task url: %s", err)
		return
	}

	requestJson := &server.TaskDoneRequest{
		TaskName:   task.Name,
		TaskStatus: task.Status,
		TaskResult: task.Detail.Result,
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
}

func verifyCommand(command string) bool {
	// todo: some sensitive commands
	return true
}
