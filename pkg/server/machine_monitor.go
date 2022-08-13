package server

import (
	"math/rand"
	"time"
)

func (s *Server) StartMachineMonitor() {
	for range time.Tick(time.Duration(rand.Intn(10)) * time.Second) {
		s.startMachineMonitorCheck()
	}
}

func (s *Server) startMachineMonitorCheck() {
	store := GetMachineStore()
	for k, each := range store.Items() {
		if !each.IsAlive() {
			store.Remove(k)
			each.Stop()
		}
	}
}
