package server

type Task struct {
	Name           string         `json:"name"`
	MachinePattern string         `json:"machinePattern"`
	Status         TaskStatus     `json:"status"`
	Detail         TaskDetail     `json:"detail"`
	Assignees      []string       `json:"assignees"`
	Result         []*AgentResult `json:"result"`
}

// TaskDetail describe how it works
type TaskDetail struct {
	Command   string `json:"command"`
	Workspace string `json:"workspace"`
	Timeout   int    `json:"timeout"`
}

type AgentResult struct {
	MachineLabel string            `json:"machineLabel"`
	Status       AgentResultStatus `json:"status"`
	Msg          string            `json:"msg"`
	Output       []string          `json:"output"`
}

func (agentResult *AgentResult) Failed() bool {
	return agentResult.Status == AgentStatusError
}

type TaskQueue []*Task

func CreateNewTask() *Task {
	return &Task{
		Name:           "",
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
