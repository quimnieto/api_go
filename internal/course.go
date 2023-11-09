package mooc

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Course struct {
	id       CourseId
	name     CourseName
	duration CourseDuration
}

func NewCourse(id, name, duration string) (Course, error) {
	courseId, err := NewCourseId(id)

	if err != nil {
		return Course{}, err
	}

	courseName, err := NewCourseName(name)

	if err != nil {
		return Course{}, err
	}

	CourseDuration, err := NewCourseDuration(duration)

	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       courseId,
		name:     courseName,
		duration: CourseDuration,
	}, nil
}

func (c Course) Id() string {
	return c.id.Value
}

func (c Course) Name() string {
	return c.name.Value
}

func (c Course) Duration() string {
	return c.duration.Value
}

type CourseId struct {
	Value string
}

var ErrInvalidCourseID = errors.New("invalid Course ID")

func NewCourseId(value string) (CourseId, error) {
	v, err := uuid.Parse(value)

	if err != nil {
		return CourseId{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseId{
		Value: v.String(),
	}, nil
}

type CourseName struct {
	Value string
}

var ErrInvalidCourseName = errors.New("invalid Course Name")

func NewCourseName(value string) (CourseName, error) {
	var courseNameString = strings.TrimSpace(value)

	if len(courseNameString) < 5 {
		return CourseName{}, fmt.Errorf("%w: %s", ErrInvalidCourseName, value)
	}

	return CourseName{
		Value: courseNameString,
	}, nil
}

type CourseDuration struct {
	Value string
}

var ErrInvalidCourseDuration = errors.New("invalid Course duration")

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, fmt.Errorf("%w: %s", ErrInvalidCourseDuration, value)
	}

	return CourseDuration{
		Value: value,
	}, nil
}

type CourseRepository interface {
	Save(ctx context.Context, course Course) error
	ById(courseId CourseId) (Course, error)
}
