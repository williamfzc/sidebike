package server

type Task struct {
	Name           string     `json:"name"`
	Type           TaskType   `json:"type"`
	MachinePattern string     `json:"machinePattern"`
	Status         TaskStatus `json:"status"`
	Detail         TaskDetail `json:"detail"`
}

type TaskDetail struct {
	// todo: duplicated!
	Name      string         `json:"name"`
	Command   string         `json:"command"`
	Workspace string         `json:"workspace"`
	Timeout   int            `json:"timeout"`
	Assignees []string       `json:"assignees"`
	Result    []*AgentResult `json:"result"`
}

type AgentResult struct {
	MachineLabel string            `json:"machineLabel"`
	Status       AgentResultStatus `json:"status"`
	Msg          string            `json:"msg"`
	Output       []string          `json:"output"`
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

func CreateNewAgentResult(assignee string) *AgentResult {
	return &AgentResult{
		assignee,
		AgentStatusInit,
		"",
		nil,
	}
}

type TaskType int

const TaskTypeCmd = 0

type AgentResultStatus int

const (
	AgentStatusOk AgentResultStatus = iota
	AgentStatusInit
	AgentStatusError
)

type TaskStatus int

const (
	// TaskStatusNew user creates this task
	TaskStatusNew TaskStatus = iota
	// TaskStatusAssigned one machine reaches this task
	TaskStatusAssigned
	// TaskStatusFinished all the machines finished this task
	TaskStatusFinished
	TaskStatusError
)
