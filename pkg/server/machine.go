package server

import (
	"sync"
	"time"
)

type Machine struct {
	Label           string        `json:"label"`
	TaskQueue       *TaskQueue    `json:"taskQueue"`
	Status          MachineStatus `json:"status"`
	FirstAliveTime  time.Time     `json:"firstAliveTime"`
	LatestAliveTime time.Time     `json:"latestAliveTime"`
}

func CreateNewMachine(label string) *Machine {
	now := time.Now()
	return &Machine{
		Label:           label,
		TaskQueue:       &TaskQueue{},
		Status:          MachineStatusNew,
		FirstAliveTime:  now,
		LatestAliveTime: now,
	}
}

func (machine *Machine) UpdateTime() {
	machine.Status = MachineStatusOnline
	machine.LatestAliveTime = time.Now()
}

func (machine *Machine) IsAlive() bool {
	return time.Now().Sub(machine.LatestAliveTime) < time.Minute
}

func (machine *Machine) AppendTask(task *Task) {
	*machine.TaskQueue = append(*machine.TaskQueue, task)
}

func (machine *Machine) GetTaskCount() int {
	return len(*machine.TaskQueue)
}

func (machine *Machine) IsEmptyTaskQueue() bool {
	return machine.GetTaskCount() == 0
}

func (machine *Machine) PopHeadTask() *Task {
	if machine.IsEmptyTaskQueue() {
		return nil
	}

	var l sync.Mutex
	l.Lock()
	defer l.Unlock()

	var ret *Task
	ret, *machine.TaskQueue = (*machine.TaskQueue)[0], (*machine.TaskQueue)[1:]
	return ret
}

type MachineStatus int

const (
	MachineStatusOffline MachineStatus = -1
	MachineStatusNew     MachineStatus = 0
	MachineStatusOnline  MachineStatus = 1
)
