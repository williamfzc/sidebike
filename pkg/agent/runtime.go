package agent

import (
	"github.com/go-cmd/cmd"
	"strings"
	"time"
)

func (agent *Agent) taskWorkMonitor() {
	for {
		task := <-agent.taskTodoQueue

		fullPath := strings.Split(task.Detail.Command, " ")
		logger.Infof("run shell: %s", fullPath)
		userCmd := cmd.NewCmd("bash", append([]string{"-c"}, fullPath...)...)
		go func() {
			<-time.After(time.Duration(task.Detail.Timeout) * time.Second)
			_ = userCmd.Stop()
		}()
		status := <-userCmd.Start()
		logger.Infof("task done: %s", status)
	}
}
