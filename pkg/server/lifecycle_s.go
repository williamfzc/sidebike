package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlePing(c *gin.Context) {
	machinePath := c.Query(FieldMachineLabel)
	GetStore().Add(machinePath, nil)

	c.JSON(http.StatusOK, Response{
		StatusCode: StatusOk,
		Msg:        "pong: " + machinePath,
	})
}
