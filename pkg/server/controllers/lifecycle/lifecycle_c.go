package lifecycle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/williamfzc/sidebike/pkg/server/controllers"
	"github.com/williamfzc/sidebike/pkg/server/services/lifecycle"
	"net/http"
)

const Prefix = "/api/lifecycle"

func withPrefix(old string) string {
	return fmt.Sprintf("%s%s", Prefix, old)
}

var Ping = &controllers.Mapping{
	HttpMethod: http.MethodGet,
	Path:       withPrefix("/ping"),
	Handler:    lifecycle.Ping,
}

func BuildController(engine *gin.Engine) {
	Ping.Add2Engine(engine)
}
