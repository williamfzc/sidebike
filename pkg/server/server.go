package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
}

func (server *Server) Execute() {
	router := gin.Default()
	initRouter(router)
	err := router.Run(fmt.Sprintf(":%d", server.Port))
	if err != nil {
		fmt.Printf("failed to start server: %s", err.Error())
	}
}

func initRouter(engine *gin.Engine) {
	BuildController(engine)
}
