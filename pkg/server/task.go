package server

type Task struct {
	Name           string     `json:"name"`
	Type           TaskType   `json:"type"`
	MachinePattern string     `json:"machinePattern"`
	Status         TaskStatus `json:"status"`
	Detail         TaskDetail `json:"detail"`
}

type TaskDetail struct {
	Command   string   `json:"command"`
	Workspace string   `json:"workspace"`
	Timeout   int      `json:"timeout"`
	Result    []string `json:"result"`
	Assignee  string   `json:"assignee"`
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

type TaskType int

const TaskTypeCmd = 0

type TaskStatus int

const (
	TaskStatusNew TaskStatus = iota
	TaskStatusAssigned
	TaskStatusFinished
	TaskStatusError
)
