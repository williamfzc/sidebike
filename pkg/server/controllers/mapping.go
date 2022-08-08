package controllers

import "github.com/gin-gonic/gin"

type Mapping struct {
	HttpMethod string
	Path       string
	Handler    func(ctx *gin.Context)
}

func (mapping *Mapping) Add2Engine(engine *gin.Engine) {
	engine.Handle(mapping.HttpMethod, mapping.Path, mapping.Handler)
}
