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
	for _, each := range GetMachineStore().GetAll() {
		if !each.IsAlive() {
			each.Status = MachineStatusOffline
		}
	}
}
