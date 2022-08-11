package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*Config
}

func CreateNewServer(config *Config) *Server {
	// default values
	if config.Port == 0 {
		config.Port = 9410
	}

	return &Server{config}
}

func (s *Server) Execute() {
	// by default
	if !s.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	go s.startMachineMonitor()

	engine := gin.Default()
	InitRouter(engine)
	logger.Infof("sidebike server started")
	err := engine.Run(fmt.Sprintf(":%d", s.Port))
	if err != nil {
		fmt.Printf("failed to start server: %s", err.Error())
	}
}

func InitRouter(engine *gin.Engine) {
	rootGroup := engine.Group("/api")
	v1Group := rootGroup.Group("/v1")

	// lifecycle
	lifeCycleGroup := v1Group.Group("/lifecycle")
	Ping.Add2Engine(lifeCycleGroup)

	// task
	taskGroup := v1Group.Group("/task")
	PostTask.Add2Engine(taskGroup)
	NewTask.Add2Engine(taskGroup)
	AssignTask.Add2Engine(taskGroup)
	DoneTask.Add2Engine(taskGroup)
	QueryTask.Add2Engine(taskGroup)

	// machine
	machineGroup := v1Group.Group("/machine")
	QueryMachine.Add2Engine(machineGroup)
}
