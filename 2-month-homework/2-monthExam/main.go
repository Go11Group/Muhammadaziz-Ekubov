package main

import (
	"add/handLer"
	"add/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	user := postgres.NewUserRepo(db)
	course := postgres.NewCourseRepo(db)
	lesson := postgres.NewLessonRepo(db)
	enrollment := postgres.NewEnrollmentRepo(db)

	handler := handLer.Handler{Db: db, User: user, Course: course, Lesson: lesson, Enrollment: enrollment}
	handler.Login()
}
