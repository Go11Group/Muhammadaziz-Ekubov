package postgres

import (
	"add/model"
	"database/sql"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) Create(course *model.Course) error {
	_, err := c.DB.Exec("insert into course (id , name, field) values ($1, $2, $3)", course.Id, course.Name, course.Field)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) Update(id int, name, field string) error {
	_, err := c.DB.Exec("update course set name = $1, field = $2 where id = $3", name, field, id)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) Delete(id int, name string) error {
	_, err := c.DB.Exec("delete from course where id = $1 and name = $2", id, name)
	if err != nil {
		return err
	}
	return nil
}
