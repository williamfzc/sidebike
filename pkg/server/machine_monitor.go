package server

import "time"

func (s *Server) startMachineMonitor() {
	for range time.Tick(15 * time.Second) {
		s.startMachineMonitorCheck()
	}
}

func (s *Server) startMachineMonitorCheck() {
	for _, each := range GetMachineStore().GetAll() {
		if !each.IsAlive() {
			logger.Infof("machine %s offline", each.Label)
			each.Status = MachineStatusOffline
		}
	}
}
