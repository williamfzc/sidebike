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
	// handle some default values
	if config.MachineLabel == "" {
		config.MachineLabel = uuid.New().String()
		logger.Warnf("no machineLabel found, use random name: %s", config.MachineLabel)
	}
	if config.Period == 0 {
		config.Period = 15
	}
	if config.Registry.Address == "" {
		config.Registry.Address = "127.0.0.1"
	}
	if config.Registry.Port == 0 {
		config.Registry.Port = 9410
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
