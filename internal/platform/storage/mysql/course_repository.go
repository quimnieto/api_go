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
	fmt.Print("leeeeeel")
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
