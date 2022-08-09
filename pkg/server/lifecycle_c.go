package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Prefix = "/api/lifecycle"
const FieldMachineLabel = "path"

func withPrefix(old string) string {
	return fmt.Sprintf("%s%s", Prefix, old)
}

var Ping = &Mapping{
	HttpMethod: http.MethodGet,
	Path:       withPrefix("/ping"),
	Handler:    HandlePing,
}

func BuildController(engine *gin.Engine) {
	Ping.Add2Engine(engine)
}
