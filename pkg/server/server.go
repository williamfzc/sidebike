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

	// lifecycle
	Ping.Add2Engine(engine)
	// task
	PostTask.Add2Engine(engine)
	NewTask.Add2Engine(engine)
	AssignTask.Add2Engine(engine)
	DoneTask.Add2Engine(engine)
	QueryTask.Add2Engine(engine)

	return engine
}
