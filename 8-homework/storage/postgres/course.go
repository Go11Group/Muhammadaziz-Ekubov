package postgres

import (
	"database/sql"
	"add/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) Create(course *model.Course) error {

	return nil
}
