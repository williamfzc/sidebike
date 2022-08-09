package server

type Task struct {
	Name           string
	Type           int
	MachinePattern string
	Status         int
}

type TaskQueue []*Task

type TaskAssign struct {
	MachinePath string
}

func CreateNewTask() *Task {
	return &Task{
		Name:           "",
		Type:           TaskTypeCmd,
		MachinePattern: "",
		Status:         TaskStatusNew,
	}
}

const TaskTypeCmd = 0

const TaskStatusNew = 0
const TaskStatusAssigned = 1
const TaskStatusDoing = 2
const TaskStatusFinished = 3
const TaskStatusError = 4
