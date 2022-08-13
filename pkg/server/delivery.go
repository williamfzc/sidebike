package server

import (
	"fmt"
	"regexp"
)

func DeliverTask(newTask *Task) string {
	compiled, err := regexp.Compile(newTask.MachinePattern)
	if err != nil {
		msg := fmt.Sprintf("parse pattern error: %s", err)
		logger.Error(msg)
		return msg
	}

	// ok this task is valid, save it
	GetTaskStore().Set(newTask.Name, newTask)

	store := GetMachineStore()
	logger.Infof("matching machines: %s", compiled)
	for _, machinePath := range store.Keys() {
		logger.Info("checking machine: %s", machinePath)
		if compiled.Match([]byte(machinePath)) {
			logger.Infof("machine %s matched, append task", machinePath)
			machine, ok := store.Get(machinePath)

			if ok {
				machine.SubmitTask(newTask)
			} else {
				logger.Warnf("machine %s offline", machinePath)
			}
		}
	}
	return ""
}

func RequestTask(request *TaskAssignRequest) (*Task, string) {
	store := GetMachineStore()
	machine, ok := store.Get(request.MachineLabel)
	if !ok {
		return nil, "no machine mapping"
	}

	task := machine.PopHeadTask()
	if task == nil {
		return nil, "no task in machine queue"
	}
	task.Detail.Assignees = append(task.Detail.Assignees, machine.Label)
	task.Status = TaskStatusAssigned
	return task, ""
}
