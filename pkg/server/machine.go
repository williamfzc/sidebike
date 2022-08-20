package server

import (
	"golang.org/x/net/context"
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
	*sync.Mutex
	done *context.CancelFunc
}

func CreateNewMachine(label string) *Machine {
	ctx, cancel := context.WithCancel(context.Background())
	now := time.Now()
	machine := &Machine{
		Label:           label,
		TaskQueue:       &TaskQueue{},
		Status:          MachineStatusNew,
		FirstAliveTime:  now,
		LatestAliveTime: now,
		eventQueue:      make(chan *MachineEvent),
		done:            &cancel,
	}

	go machine.StartEventWatcher(ctx)
	go machine.StartStatusWatcher(ctx)
	return machine
}

func (machine *Machine) StartStatusWatcher(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
		logger.Infof("machine %s ticker shutdown", machine.Label)
	}()

	for {
		select {
		case <-ticker.C:
			if !machine.IsAlive() {
				logger.Infof("machine %s offline", machine.Label)
				machine.Status = MachineStatusOffline
			} else {
				machine.Status = MachineStatusOnline
			}
		case <-ctx.Done():
			return
		}
	}
}

func (machine *Machine) StartEventWatcher(ctx context.Context) {
	defer func() {
		close(machine.eventQueue)
		logger.Infof("machine %s event queue shutdown", machine.Label)
	}()

	for {
		select {
		case newEvent, running := <-machine.eventQueue:
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
		case <-ctx.Done():
			return
		}
	}
}

func (machine *Machine) Sync() {
	if machine.Status == MachineStatusOffline {
		return
	}
	machine.eventQueue <- &MachineEvent{MachineEventSync, nil}
}

func (machine *Machine) SubmitTask(task *Task) {
	if machine.Status == MachineStatusOffline {
		return
	}
	machine.eventQueue <- &MachineEvent{MachineEventNewTask, task}
}

func (machine *Machine) Stop() {
	machine.Lock()
	defer machine.Unlock()

	(*machine.done)()
	logger.Infof("unregister machine %s because of offline", machine.Label)
}

func (machine *Machine) PopHeadTask() *Task {
	machine.Lock()
	defer machine.Unlock()

	if machine.IsEmptyTaskQueue() {
		return nil
	}
	if !machine.IsAlive() {
		return nil
	}

	var ret *Task
	realQueue := machine.TaskQueue
	ret, *realQueue = (*realQueue)[0], (*realQueue)[1:]
	return ret
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

func (machine *Machine) appendTask(task *Task) {
	*machine.TaskQueue = append(*machine.TaskQueue, task)
}

func (machine *Machine) updateTime() {
	machine.Status = MachineStatusOnline
	machine.LatestAliveTime = time.Now()
}
