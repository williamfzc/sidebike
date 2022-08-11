package server

type TaskAssignRequest struct {
	MachineLabel string
}

type TaskDoneRequest struct {
	TaskName string
	Result   *AgentResult
}
