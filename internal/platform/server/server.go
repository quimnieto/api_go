package server

import (
	"api_go/internal/create"
	"api_go/internal/platform/server/handler/courses"
	"api_go/internal/platform/server/handler/health"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string

	// dependencies
	courseCreator create.CourseCreator
}

func New(host string, port uint, courseCreator create.CourseCreator) Server {
	server := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		courseCreator: courseCreator,
	}

	server.registerRoutes()

	return server
}

func (server *Server) Run() error {
	log.Println("Server running on", server.httpAddr)
	return server.engine.Run(server.httpAddr)
}

func (server *Server) registerRoutes() {
	server.engine.GET("/health", health.CheckHandler())
	server.engine.POST("/courses", courses.CreateCourseHandler(server.courseCreator))
}
