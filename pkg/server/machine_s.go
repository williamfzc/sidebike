package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleQueryMachine(c *gin.Context) {
	machines := GetMachineStore().Items()
	c.JSON(http.StatusOK, Response{Signal: SignalOk, Data: machines})
}
