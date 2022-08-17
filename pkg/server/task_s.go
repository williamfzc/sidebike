package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

	// https://stackoverflow.com/a/20234207
	newTask.Name = fmt.Sprintf("%s-%s", newTask.Name, time.Now().Format("20060102150405"))

	errMsg := DeliverTask(newTask)
	if errMsg != "" {
		c.JSON(http.StatusBadRequest, Response{
			Signal: SignalError,
			Msg:    errMsg,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Signal: SignalOk,
		Msg:    "task registered",
		Data:   &newTask,
	})
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

	task, errMsg := RequestTask(taskAssignRequest)
	if errMsg != "" {
		c.JSON(http.StatusOK, Response{Signal: SignalError, Msg: errMsg})
		return
	}
	c.JSON(http.StatusOK, Response{Signal: SignalOk, Data: task})
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
	task, ok := GetTaskStore().Get(taskDoneRequest.TaskName)
	if ok {
		agentResult := taskDoneRequest.Result
		logger.Infof("task %v, agent %v, result: %v",
			task.Name,
			agentResult.MachineLabel,
			agentResult.Status,
		)
		task.Result = append(task.Result, agentResult)
	}
	c.JSON(http.StatusOK, Response{Signal: SignalOk})
}

func HandleQueryTask(c *gin.Context) {
	taskPrefix := c.Query(FieldTaskPrefix)
	tasks := GetTaskStore().Items()
	if taskPrefix == "" {
		c.JSON(http.StatusOK, Response{Signal: SignalOk, Data: tasks})
		return
	}

	var tasksAfterFilter []*Task
	for i := range tasks {
		item := tasks[i]
		if strings.HasPrefix(item.Name, taskPrefix) {
			tasksAfterFilter = append(tasksAfterFilter, &item)
		}
	}
	c.JSON(http.StatusOK, Response{Signal: SignalOk, Data: tasksAfterFilter})
}
