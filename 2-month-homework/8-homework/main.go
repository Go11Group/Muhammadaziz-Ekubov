package main

import (
	"add/model"
	"add/storage/postgres"
	"fmt"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	st := postgres.NewStudentRepo(db)

	users, err := st.GetAllStudents()
	if err != nil {
		panic(err)
	}

	user, _ := st.GetByID(users[2].ID)

	fmt.Println(users)

	fmt.Println(user)

	cr := postgres.NewCourseRepo(db)
	_ = cr.Create(&model.Course{})
}
