package server

import "sync"

type Machine struct {
	Label     string
	TaskQueue *TaskQueue
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
