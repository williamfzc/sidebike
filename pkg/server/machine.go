package server

import (
	"sync"
	"time"
)

type MachineStatus int

const (
	MachineStatusOffline MachineStatus = -1
	MachineStatusNew     MachineStatus = 0
	MachineStatusOnline  MachineStatus = 1
)

type MachineEvent struct {
	Flag MachineEventFlag
	Data interface{}
}

type MachineEventFlag int

const (
	MachineEventNewTask MachineEventFlag = iota
	MachineEventSync
)

type Machine struct {
	Label           string        `json:"label"`
	TaskQueue       *TaskQueue    `json:"taskQueue"`
	Status          MachineStatus `json:"status"`
	FirstAliveTime  time.Time     `json:"firstAliveTime"`
	LatestAliveTime time.Time     `json:"latestAliveTime"`
	eventQueue      chan *MachineEvent
}

func CreateNewMachine(label string) *Machine {
	now := time.Now()
	machine := &Machine{
		Label:           label,
		TaskQueue:       &TaskQueue{},
		Status:          MachineStatusNew,
		FirstAliveTime:  now,
		LatestAliveTime: now,
		eventQueue:      make(chan *MachineEvent),
	}
	go machine.StartEventWatcher()
	go machine.StartStatusWatcher()
	return machine
}

func (machine *Machine) StartStatusWatcher() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		if !machine.IsAlive() {
			logger.Infof("machine %s offline", machine.Label)
			machine.Status = MachineStatusOffline
		}
	}
}

func (machine *Machine) StartEventWatcher() {
	for {
		newEvent, running := <-machine.eventQueue
		if !running {
			logger.Infof("machine watcher shut down: %s", machine.Label)
			return
		}

		switch newEvent.Flag {
		case MachineEventSync:
			machine.updateTime()

		case MachineEventNewTask:
			task, ok := newEvent.Data.(*Task)
			if ok {
				machine.appendTask(task)
			}
		}
	}
}

func (machine *Machine) Sync() {
	machine.eventQueue <- &MachineEvent{MachineEventSync, nil}
}

func (machine *Machine) SubmitTask(task *Task) {
	machine.eventQueue <- &MachineEvent{MachineEventNewTask, task}
}

func (machine *Machine) Stop() {
	// todo: improve db first
	//close(machine.eventQueue)
}

func (machine *Machine) IsAlive() bool {
	return time.Now().Sub(machine.LatestAliveTime) < time.Minute
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
	realQueue := machine.TaskQueue
	ret, *realQueue = (*realQueue)[0], (*realQueue)[1:]
	return ret
}

func (machine *Machine) appendTask(task *Task) {
	*machine.TaskQueue = append(*machine.TaskQueue, task)
}

func (machine *Machine) updateTime() {
	machine.Status = MachineStatusOnline
	machine.LatestAliveTime = time.Now()
}
