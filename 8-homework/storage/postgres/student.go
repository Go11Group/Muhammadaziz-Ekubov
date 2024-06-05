package postgres

import (
	"add/model"
	"database/sql"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
	query := "SELECT * FROM students"
	rows, err := u.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []model.User{}
	for rows.Next() {
		student := model.User{}
		err := rows.Scan(&student.ID, &student.Name, &student.Course, student.Gender, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (u *StudentRepo) Insert(id, name, gender, course string, age int) error {
	query := "INSERT INTO students (id, name, age, course, gender) VALUES ($1, $2, $3, $4, $5), id, name, age, course, gender"

	_, err := u.Db.Exec(query, id, name, age, gender, course)
	if err != nil {
		return err
	}
	return nil
}

func (u *StudentRepo) Update(id, name, gender, course string, age int) error {
	_, err := u.Db.Exec("Update student Set where id = $1, name = $2, gender = $4, course = $5, age = $3", id, name, gender, course, age)
	if err != nil {
		return err
	}
	return nil
}

func (u *StudentRepo) Delete(id, name, gender string) error {
	_, err := u.Db.Exec("Delete from students where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
