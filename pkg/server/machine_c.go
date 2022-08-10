package server

import (
	"fmt"
	"net/http"
)

const PrefixMachine = "/api/machine"

func withMachinePrefix(old string) string {
	return fmt.Sprintf("%s%s", PrefixMachine, old)
}

var QueryMachine = &Mapping{
	HttpMethod: http.MethodGet,
	Path:       withMachinePrefix("/"),
	Handler:    HandleQueryMachine,
}
