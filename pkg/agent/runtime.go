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
	agentResult := server.CreateNewAgentResult(agent.MachineLabel)
	ok := verifyCommand(commandToRun)
	if !ok {
		agentResult.Status = server.AgentStatusError
		agentResult.Msg = "command verify error"
		agent.updateTaskStatus(task, agentResult)
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
	logger.Infof("taskDetail done: %v", status)

	if status.Exit == 0 {
		agentResult.Status = server.AgentStatusOk
	} else {
		agentResult.Status = server.AgentStatusError
	}

	stdoutRet := userCmd.Status().Stdout
	if len(stdoutRet) > ResultLineLimit {
		stdoutRet = stdoutRet[ResultLineLimit:]
	}
	agentResult.Output = stdoutRet
	agent.updateTaskStatus(task, agentResult)
}

func (agent *Agent) updateTaskStatus(task *server.Task, agentResult *server.AgentResult) {
	// updateTaskStatus
	finalUrl, err := agent.GetUrlTaskDone()
	if err != nil {
		logger.Errorf("failed to gen task url: %s", err)
		return
	}

	requestJson := &server.TaskDoneRequest{
		TaskName: task.Name,
		Result:   agentResult,
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
