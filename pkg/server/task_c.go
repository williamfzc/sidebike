package server

import (
	"net/http"
)

var PostTask = &Mapping{
	HttpMethod: http.MethodPost,
	Path:       "/",
	Handler:    HandlePostTask,
}

var NewTask = &Mapping{
	HttpMethod: http.MethodPost,
	Path:       "/status/new",
	Handler:    HandlePostTask,
}

var AssignTask = &Mapping{
	HttpMethod: http.MethodPost,
	Path:       "/status/assigned",
	Handler:    HandleAssignTask,
}

var DoneTask = &Mapping{
	HttpMethod: http.MethodPost,
	Path:       "/status/done",
	Handler:    HandleDoneTask,
}

var QueryTask = &Mapping{
	HttpMethod: http.MethodGet,
	Path:       "/",
	Handler:    HandleQueryTask,
}
