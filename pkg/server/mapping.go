package server

import "github.com/gin-gonic/gin"

type Mapping struct {
	HttpMethod string
	Path       string
	Handler    func(ctx *gin.Context)
}

func (mapping *Mapping) Add2Engine(engine *gin.RouterGroup) {
	engine.Handle(mapping.HttpMethod, mapping.Path, mapping.Handler)
	// update this path for client
	mapping.Path = engine.BasePath() + mapping.Path
}
