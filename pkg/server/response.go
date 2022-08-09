package server

type Response struct {
	StatusCode int    `json:"code"`
	Msg        string `json:"msg"`
	Data       string `json:"data"`
}

const (
	StatusOk    = 0
	StatusError = 1
)
