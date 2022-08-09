package server

type Response struct {
	Signal int         `json:"signal"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type TaskResponse struct {
	*Response
	Data Task `json:"data"`
}

// these signs will tell agents about what they should do
const (
	// SignalError something wrong, check the msg in resp
	SignalError = -1
	// SignalOk need nothing to do
	SignalOk = 0
	// SignalNewTask new task arrived
	SignalNewTask = 1
)
