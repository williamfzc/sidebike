package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	Port int
}

func (server *Server) Execute() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	fmt.Printf("server port: %d", server.Port)
	err := router.Run(fmt.Sprintf(":%d", server.Port))
	if err != nil {
		fmt.Printf("failed to start server: %s", err.Error())
	}
}
