package server

type Task struct {
	Name           string     `json:"name"`
	Type           int        `json:"type"`
	MachinePattern string     `json:"machinePattern"`
	Status         int        `json:"status"`
	Detail         TaskDetail `json:"detail"`
}

type TaskDetail struct {
	Command string   `json:"command"`
	Timeout int      `json:"timeout"`
	Result  []string `json:"result"`
}

type TaskQueue []*Task

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
const TaskStatusFinished = 2
const TaskStatusError = 3
