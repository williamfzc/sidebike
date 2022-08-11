package server

import (
	"math/rand"
	"time"
)

func (s *Server) StartTaskMonitor() {
	for range time.Tick(time.Duration(rand.Intn(10)) * time.Second) {
		s.startTaskMonitorCheck()
	}
}

func (s *Server) startTaskMonitorCheck() {
	for _, each := range GetTaskStore().GetAll() {
		assigneeNum := len(each.Detail.Assignees)

		// all the results collected
		if (assigneeNum != 0) && assigneeNum == len(each.Detail.Result) {
			for _, eachResult := range each.Detail.Result {
				if eachResult.Failed() {
					each.Status = TaskStatusError
					break
				}
			}
			each.Status = TaskStatusFinished
		}
	}
}
