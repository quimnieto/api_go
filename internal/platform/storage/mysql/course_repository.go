package mysql

import (
	mooc "api_go/internal"
	"context"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (repository *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))

	query, args := courseSQLStruct.InsertInto(table, sqlCourse{
		ID:       course.Id(),
		Name:     course.Name(),
		Duration: course.Duration(),
	}).Build()

	_, err := repository.db.ExecContext(ctx, query, args...)

	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

func (repository *CourseRepository) ById(courseId mooc.CourseId) (mooc.Course, error) {
	var courseSQL sqlCourse

	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))

	query, args := courseSQLStruct.SelectFrom(table).Where("id = ?", courseId.Value).Build()

	err := repository.db.QueryRowContext(ctx, query, args...).Scan(
		&courseSQL.ID,
		&courseSQL.Name,
		&courseSQL.Duration,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return mooc.Course{}, fmt.Errorf("course not found")
		}

		return mooc.Course{}, fmt.Errorf("error trying to retrieve course from database: %v", err)
	}

	course, err := mooc.NewCourse(
		courseSQL.ID,
		courseSQL.Name,
		courseSQL.Duration,
	)

	if err != nil {
		return mooc.Course{}, fmt.Errorf("Error trying to build the course")
	}

	return course, nil
}
