package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s: %d", host, port),
	}
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running ON", s.httpAddr)
	s.engine.SetTrustedProxies([]string{"192.168.1.3"})
	return s.engine.Run(s.httpAddr)
}

func (s *Server) RegisterRoute(httpMethod string, relativePath string, handler gin.HandlerFunc) {
	s.engine.RouterGroup.Handle(httpMethod, relativePath, handler)
}
