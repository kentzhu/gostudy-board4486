package httpserver

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	engine *gin.Engine
}

func CreatServer() *Server {
	server := Server{
		engine: gin.Default(),
	}
	return &server
}

func (server *Server) Run() {
	err := server.engine.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) BindRouteRule(ruleBinder func(engine *gin.Engine)) {
	ruleBinder(server.engine)
}
