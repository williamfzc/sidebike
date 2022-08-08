package agent

import (
	"fmt"
	"time"
)

type Agent struct {
	Conf *Config
}

func (agent *Agent) GetPeriod() time.Duration {
	return time.Duration(agent.Conf.Period) * time.Second
}

func (agent *Agent) GetRegistry() string {
	return fmt.Sprintf("%s:%d", agent.Conf.Registry.Address, agent.Conf.Registry.Port)
}

func (agent *Agent) Run() {
	period := agent.GetPeriod()
	for range time.Tick(period) {
		fmt.Printf("exec agent: %s\n", agent.GetRegistry())
	}
	select {}
}
