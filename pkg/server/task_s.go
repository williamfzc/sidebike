package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

const FieldTaskPrefix = "taskPrefix"

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
			machine.SubmitTask(newTask)
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
	machine, ok := store.GetWithType(taskAssignRequest.MachineLabel)
	if !ok {
		c.JSON(SignalOk, Response{Signal: SignalError, Msg: "no machine mapping"})
		return
	}

	task := machine.PopHeadTask()
	if task != nil {
		task.Detail.Assignees = append(task.Detail.Assignees, machine.Label)
		task.Status = TaskStatusAssigned
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
		agentResult := taskDoneRequest.Result
		logger.Infof("task %v, agent %v, result: %v",
			task.Name,
			agentResult.MachineLabel,
			agentResult.Status,
		)
		task.Detail.Result = append(task.Detail.Result, agentResult)
	}
	c.JSON(http.StatusOK, Response{Signal: SignalOk})
}

func HandleQueryTask(c *gin.Context) {
	taskPrefix := c.Query(FieldTaskPrefix)
	tasks := GetTaskStore().GetAll()
	if taskPrefix == "" {
		c.JSON(http.StatusOK, Response{Signal: SignalOk, Data: tasks})
		return
	}

	var tasksAfterFilter []*Task
	for i := range tasks {
		item := tasks[i]
		if strings.HasPrefix(item.Name, taskPrefix) {
			tasksAfterFilter = append(tasksAfterFilter, item)
		}
	}
	c.JSON(http.StatusOK, Response{Signal: SignalOk, Data: tasksAfterFilter})
}
