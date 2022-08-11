package server

import (
	"net/http"
)

const FieldMachineLabel = "path"

var Ping = &Mapping{
	HttpMethod: http.MethodGet,
	Path:       "/ping",
	Handler:    HandlePing,
}
