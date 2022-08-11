package server

import (
	"net/http"
)

var QueryMachine = &Mapping{
	HttpMethod: http.MethodGet,
	Path:       "/",
	Handler:    HandleQueryMachine,
}
