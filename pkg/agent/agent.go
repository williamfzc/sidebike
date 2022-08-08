package agent

import (
	"fmt"
	"github.com/williamfzc/sidebike/pkg/server/controllers/lifecycle"
	"net/http"
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

func (agent *Agent) GetUrlPing() string {
	return "http://" + agent.GetRegistry() + lifecycle.Ping.Path
}

func (agent *Agent) Run() {
	period := agent.GetPeriod()
	for range time.Tick(period) {
		resp, err := http.Get(agent.GetUrlPing())
		if err != nil {
			fmt.Printf("request error: %s\n", err)
		} else {
			fmt.Printf("ping backend: %s\n", resp)
		}
	}
	select {}
}
