package agent

type Agent struct {
	*Config
}

func (agent *Agent) Run() {
	go agent.heartBeat()
	select {}
}
