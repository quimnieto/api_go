package server

import (
	"api_go/internal/platform/server/handler/courses"
	"api_go/internal/platform/server/handler/health"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateCourseHandler())
}
