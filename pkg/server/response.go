package server

type Response struct {
	StatusCode int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

type TaskResponse struct {
	*Response
	Data Task `json:"data"`
}

const (
	StatusError   = -1
	StatusOk      = 0
	StatusSync    = 1
	StatusNewTask = 2
)
