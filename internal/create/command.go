package create

import (
	"api_go/kit/command"
	"context"
	"errors"
)

const CreateCourseCommandType command.Type = "command.create.course"

type CreateCourseCommand struct {
	id       string
	name     string
	duration string
}

func NewCreateCourseCommand(id, name, duration string) CreateCourseCommand {
	return CreateCourseCommand{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (command CreateCourseCommand) Type() command.Type {
	return CreateCourseCommandType
}

type CreateCourseCommandHandler struct {
	courseCreator CourseCreator
}

func NewCreateCourseCommandHandler(courseCreator CourseCreator) CreateCourseCommandHandler {
	return CreateCourseCommandHandler{
		courseCreator: courseCreator,
	}
}

func (handler CreateCourseCommandHandler) Handle(ctx context.Context, command command.Command) error {
	createCourseCommand, ok := command.(CreateCourseCommand)

	if !ok {
		return errors.New("unexpected command")
	}

	return handler.courseCreator.CreateCourse(ctx, createCourseCommand.id, createCourseCommand.name, createCourseCommand.duration)
}
