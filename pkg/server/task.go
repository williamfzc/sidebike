package server

type Task struct {
	Name           string     `json:"name"`
	Type           int        `json:"type"`
	MachinePattern string     `json:"machinePattern"`
	Status         int        `json:"status"`
	Detail         TaskDetail `json:"detail"`
}

type TaskDetail struct {
	Command string `json:"command"`
	Timeout int    `json:"timeout"`
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

// todo
const TaskStatusNew = 0
const TaskStatusAssigned = 1
const TaskStatusDoing = 2
const TaskStatusFinished = 3
const TaskStatusError = 4
