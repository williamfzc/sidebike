package server

type TaskAssignRequest struct {
	MachineLabel string
}

type TaskDoneRequest struct {
	TaskName   string
	TaskStatus TaskStatus
	TaskResult []string
}
