package agent

type Agent struct {
	taskRequestQueue chan TaskRequestEvent
	taskWorkQueue    chan TaskWorkEvent
	*Config
}

type Event struct{}
type TaskRequestEvent Event
type TaskWorkEvent Event

func CreateAgent(config *Config) *Agent {
	return &Agent{make(chan TaskRequestEvent), make(chan TaskWorkEvent), config}
}

func (agent *Agent) Run() {
	go agent.heartBeat()
	go agent.taskRequestMonitor()
	select {}
}
