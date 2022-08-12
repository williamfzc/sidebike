package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlePing(c *gin.Context) {
	machinePath := c.Query(FieldMachineLabel)

	// update machine store
	store := GetMachineStore()
	if !store.Has(machinePath) {
		machine := CreateNewMachine(machinePath)
		store.Set(machinePath, machine)
	} else {
		if machine, ok := store.Get(machinePath); ok {
			machine.Sync()
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
