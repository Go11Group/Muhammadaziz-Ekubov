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
	rows, err := u.Db.Query(`select s.id, s.name, age, gender, c.name from student s
					left join course c on c.id = s.course_id `)
	if err != nil {
		return nil, err
	}

	var users []model.User
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByID(id string) (model.User, error) {
	var user model.User
	//
	//err := u.Db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
	//				left join course c on c.id = s.course_id where s.id = $1`, id).
	//	Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	//if err != nil {
	//	return model.User{}, err
	//}

	return user, nil
}

func Create(db *sql.DB, user model.User) error {
	//
	//query := `
	//    INSERT INTO users (id, name, age, gender, course)
	//    VALUES (default, 'Johan', 32, 'm', 'golang')
	//`
	//stmt, err := db.Prepare(query)
	//if err != nil {
	//	return fmt.Errorf("failed to prepare insert statement: %v", err)
	//}
	//defer stmt.Close()
	//

	//_, err = stmt.Exec(user.Name, user.Email, user.Age)
	//if err != nil {

	return nil
}

func Update(db *sql.DB, user model.User) error {

	//err := db.QueryRow(`update student Set age = $1 where id = $2`, `user.Age, user.ID`)
	//	if err != nil {
	//		return nil
	//	}

	return nil
}

func Delete(db *sql.DB, id string) error {

	//query := `
	//    DELETE FROM users
	//    WHERE id = $1
	//`
	//stmt, err := db.Prepare(query)
	//if err != nil {
	//	return fmt.Errorf("failed to prepare delete statement: %v", err)
	//}
	//defer stmt.Close()
	//

	//_, err = stmt.Exec(id)
	//if err != nil {
	//	return fmt.Errorf("failed to execute delete statement: %v", err)
	//}

	return nil
}

