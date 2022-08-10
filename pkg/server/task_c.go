package server

import (
	"fmt"
	"net/http"
)

const PrefixTask = "/api/task"

func withTaskPrefix(old string) string {
	return fmt.Sprintf("%s%s", PrefixTask, old)
}

var PostTask = &Mapping{
	HttpMethod: http.MethodPost,
	Path:       withTaskPrefix("/"),
	Handler:    HandlePostTask,
}

var NewTask = &Mapping{
	HttpMethod: http.MethodPost,
	Path:       withTaskPrefix("/status/new"),
	Handler:    HandlePostTask,
}

var AssignTask = &Mapping{
	HttpMethod: http.MethodPost,
	Path:       withTaskPrefix("/status/assigned"),
	Handler:    HandleAssignTask,
}

var DoneTask = &Mapping{
	HttpMethod: http.MethodPost,
	Path:       withTaskPrefix("/status/done"),
	Handler:    HandleDoneTask,
}

var QueryTask = &Mapping{
	HttpMethod: http.MethodGet,
	Path:       withTaskPrefix("/"),
	Handler:    HandleQueryTask,
}
