package agent

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-cmd/cmd"
	"github.com/williamfzc/sidebike/pkg/server"
)

const ResultLineLimit = 30

func (agent *Agent) StartTaskWorkMonitor() {
	for {
		task := <-agent.taskTodoQueue
		agent.triggerTaskWork(task)
	}
}

func (agent *Agent) triggerTaskWork(task *server.Task) {
	commandToRun := task.Detail.Command
	ok := verifyCommand(commandToRun)
	if !ok {
		task.Status = server.TaskStatusError
		task.Detail.Result = []string{"command verify error"}
		agent.updateTaskStatus(task)
		return
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
	agent.updateTaskStatus(task)
}

func (agent *Agent) updateTaskStatus(task *server.Task) {
	// updateTaskStatus
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
