package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlePing(c *gin.Context) {
	machinePath := c.Query(FieldMachineLabel)

	// update machine store
	store := GetMachineStore()
	if !store.Contains(machinePath) {
		machine := &Machine{machinePath, &TaskQueue{}}
		store.Add(machinePath, machine)
	} else {
		if machine, ok := store.GetWithType(machinePath); ok {
			if !machine.IsEmptyTaskQueue() {
				c.JSON(http.StatusOK, Response{
					Signal: SignalNewTask,
				})
				return
			}
		}
	}

	// normal
	c.JSON(http.StatusOK, Response{
		Signal: SignalOk,
		Msg:    "pong: " + machinePath,
	})
}
