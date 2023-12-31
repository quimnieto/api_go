package server

import (
	"api_go/internal/platform/server/handler/courses"
	"api_go/internal/platform/server/handler/health"
	"api_go/kit/command"
	"api_go/kit/query"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	httpAddr string

	// dependencies
	commandBus command.CommandBus
	queryBus   query.QueryBus
}

func New(host string, port uint, commandBus command.CommandBus, queryBus query.QueryBus) Server {
	server := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		commandBus: commandBus,
		queryBus:   queryBus,
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
	server.engine.POST("/courses", courses.CreateCourseHandler(server.commandBus))
}
