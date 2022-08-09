package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func HandlePostTask(c *gin.Context) {
	newTask := CreateNewTask()
	err := c.BindJSON(newTask)
	if err != nil {
		msg := fmt.Sprintf("parse task error: %s", err)
		logger.Error(msg)
		c.JSON(http.StatusBadRequest, Response{
			Signal: SignalError,
			Msg:    msg,
		})
		return
	}

	compiled, err := regexp.Compile(newTask.MachinePattern)
	if err != nil {
		msg := fmt.Sprintf("parse pattern error: %s", err)
		logger.Error(msg)
		c.JSON(http.StatusBadRequest, Response{
			Signal: SignalError,
			Msg:    msg,
		})
		return
	}

	// ok this task is valid, save it
	GetTaskStore().Add(newTask.Name, newTask)

	store := GetMachineStore()
	logger.Debugf("matching machines: %s", compiled)
	for _, machinePath := range store.Keys() {
		logger.Debugf("checking machine: %s", machinePath)
		if compiled.Match([]byte(machinePath.(string))) {
			logger.Debugf("machine %s matched, append task", machinePath)
			machine, _ := store.GetWithType(machinePath)
			machine.AppendTask(newTask)
		}
	}
	c.JSON(http.StatusOK, Response{Signal: SignalOk})
}

func HandleAssignTask(c *gin.Context) {
	taskAssignRequest := &TaskAssignRequest{}
	err := c.BindJSON(taskAssignRequest)
	if err != nil {
		msg := fmt.Sprintf("parse assign error: %s", err)
		logger.Error(msg)
		c.JSON(http.StatusBadRequest, Response{
			Signal: SignalError,
			Msg:    msg,
		})
		return
	}

	store := GetMachineStore()
	machine, ok := store.GetWithType(taskAssignRequest.MachinePath)
	if !ok {
		c.JSON(SignalOk, Response{Signal: SignalError, Msg: "no machine mapping"})
		return
	}

	task := machine.PopHeadTask()
	if task != nil {
		c.JSON(SignalOk, Response{Signal: SignalOk, Data: task})
		return
	}

	// default
	c.JSON(http.StatusOK, Response{Signal: SignalOk, Msg: "no task need to run"})
}

func HandleDoneTask(c *gin.Context) {
	taskDoneRequest := &TaskDoneRequest{}
	err := c.BindJSON(taskDoneRequest)
	if err != nil {
		msg := fmt.Sprintf("parse task request error: %s", err)
		logger.Error(msg)
		c.JSON(http.StatusBadRequest, Response{
			Signal: SignalError,
			Msg:    msg,
		})
		return
	}

	// todo: name will conflict
	task, ok := GetTaskStore().GetWithType(taskDoneRequest.TaskName)
	if ok {
		logger.Infof("update task status to: %d", taskDoneRequest.TaskStatus)
		logger.Infof("task output: %s", taskDoneRequest.TaskResult)
		task.Status = taskDoneRequest.TaskStatus
		task.Detail.Result = taskDoneRequest.TaskResult
	}
	c.JSON(http.StatusOK, Response{Signal: SignalOk})
}
