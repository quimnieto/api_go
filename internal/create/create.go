package create

import (
	mooc "api_go/internal"
	"context"
)

type CourseCreator struct {
	courseRepository mooc.CourseRepository
}

func NewCourseCreator(courseRepository mooc.CourseRepository) CourseCreator {
	return CourseCreator{
		courseRepository: courseRepository,
	}
}

func (courseCreator CourseCreator) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)

	if err != nil {
		return err
	}

	return courseCreator.courseRepository.Save(ctx, course)
}
