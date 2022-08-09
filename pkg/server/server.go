package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
}

func (server *Server) Execute() {
	router := initRouter()
	err := router.Run(fmt.Sprintf(":%d", server.Port))
	if err != nil {
		fmt.Printf("failed to start server: %s", err.Error())
	}
}

func initRouter() *gin.Engine {
	engine := gin.Default()
	Ping.Add2Engine(engine)
	PostTask.Add2Engine(engine)
	AssignTask.Add2Engine(engine)
	return engine
}
