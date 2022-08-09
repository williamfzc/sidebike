package agent

import "github.com/williamfzc/sidebike/pkg/server"

type Agent struct {
	taskRequestQueue chan *Event
	taskTodoQueue    chan *server.Task
	*Config
}

type Event struct{}

func CreateAgent(config *Config) *Agent {
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
	select {}
}
