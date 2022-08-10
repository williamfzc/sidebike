package agent

import (
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
	if config.MachineLabel == "" {
		config.MachineLabel = uuid.New().String()
		logger.Warnf("no machineLabel found, use random name: %s", config.MachineLabel)
	}

	return &Agent{
		make(chan *Event),
		make(chan *server.Task),
		config,
	}
}

func (agent *Agent) Run() {
	go agent.taskWorkMonitor()
	go agent.taskRequestMonitor()
	go agent.heartBeat()
	logger.Infof("sidebike agent started")
	select {}
}
