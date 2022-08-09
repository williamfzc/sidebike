package server

import (
	"fmt"
	"net/http"
)

const PrefixLifecycle = "/api/lifecycle"
const FieldMachineLabel = "path"

func withLifecyclePrefix(old string) string {
	return fmt.Sprintf("%s%s", PrefixLifecycle, old)
}

var Ping = &Mapping{
	HttpMethod: http.MethodGet,
	Path:       withLifecyclePrefix("/ping"),
	Handler:    HandlePing,
}
