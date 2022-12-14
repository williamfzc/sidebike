package agent

import (
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/williamfzc/sidebike/pkg/server"
)

type Agent struct {
	taskRequestQueue chan *Event
	taskTodoQueue    chan *server.Task
	*Config
}

type Event struct{}

func CreateAgent(config *Config) *Agent {
	// handle some default values
	if config.MachineLabel == "" {
		config.MachineLabel = uuid.New().String()
		logger.Warnf("no machineLabel found, use random name: %s", config.MachineLabel)
	}
	if config.Period == 0 {
		config.Period = 10
	}
	if config.Registry.Address == "" {
		config.Registry.Address = "127.0.0.1"
		logger.Warn("registry address set to localhost")
	}
	if config.Registry.Port == 0 {
		config.Registry.Port = 9410
	}

	// init router for talking to backend
	gin.SetMode(gin.ReleaseMode)
	server.InitRouter(gin.New())

	return &Agent{
		make(chan *Event),
		make(chan *server.Task),
		config,
	}
}

func (agent *Agent) Run() {
	// pre check
	if !preCheck() {
		return
	}

	go agent.StartTaskWorkMonitor()
	go agent.StartTaskRequestMonitor()
	go agent.StartHeartBeatMonitor()
	logger.Infof("sidebike agent started")
	select {}
}

func preCheck() bool {
	current, err := user.Current()
	if err != nil {
		logger.Errorf("pre check error: %s", err)
		return false
	}
	if current.Username == "root" {
		logger.Errorf("SHOULD NOT <START AGENT WITH ROOT> BECAUSE OF SECURITY ISSUES")
	}
	return true
}
