package server

type TaskAssignRequest struct {
	MachinePath string
}

type TaskDoneRequest struct {
	TaskName   string
	TaskStatus int
	TaskResult []string
}
