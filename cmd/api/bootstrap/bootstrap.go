package bootstrap

import (
	"api_go/internal/create"
	"api_go/internal/platform/bus"
	"api_go/internal/platform/server"
	"api_go/internal/platform/storage/mysql"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "0.0.0.0"
	port = 8080

	dbUser = "root"
	dbPass = ""
	dbHost = "mysql"
	dbPort = "3306"
	dbName = "api_go_db"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		return err
	}

	commandBus := bus.NewCommandBus()
	queryBus := bus.NewQueryBus()

	courseRepository := mysql.NewCourseRepository(db)
	courseCreator := create.NewCourseCreator(courseRepository)
	createCourseCommandHandler := create.NewCreateCourseCommandHandler(courseCreator)
	
	commandBus.Register(create.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, commandBus, queryBus)

	return srv.Run()
}
