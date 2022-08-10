package agent

import (
	"fmt"
	"time"
)

type RegistryConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Config struct {
	Registry     RegistryConfig `json:"registry"`
	Period       int            `json:"period"`
	MachineLabel string         `json:"machineLabel"`
}

func (c *Config) GetPeriod() time.Duration {
	return time.Duration(c.Period) * time.Second
}

func (c *Config) GetRegistry() string {
	return fmt.Sprintf("%s:%d", c.Registry.Address, c.Registry.Port)
}
