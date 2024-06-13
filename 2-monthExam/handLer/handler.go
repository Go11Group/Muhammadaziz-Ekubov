package handLer

import (
	"add/storage/postgres"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Db         *sql.DB
	User       *postgres.UserRepo
	Course     *postgres.CourseRepo
	Lesson     *postgres.LessonRepo
	Enrollment *postgres.EnrollmentRepo
}

func (h *Handler) Login() {
	router := gin.Default()

	// Users API
	user := router.Group("/user")
	user.POST("/create", h.CreateUser)
	user.GET("/read/:id", h.ReadUser)
	user.PUT("/update", h.UpdateUser)
	user.DELETE("/delete/:id", h.DeleteUser)
	user.GET("/list", h.ListUser)

	// course API
	course := router.Group("/course")
	course.POST("/create", h.CreateCourse)
	course.GET("/read/:id", h.ReadCoursesByID)
	course.PUT("/update", h.UpdateCourse)
	course.DELETE("/delete/:id", h.DeleteCourseByID)
	course.GET("/list", h.listCourses)

	// lessons API
	lesson := router.Group("/lesson")
	lesson.POST("/create", h.CreateLesson)
	lesson.GET("/read/:id", h.ReadLessonsByID)
	lesson.PUT("/update/:id", h.UpdateLesson)
	lesson.DELETE("/delete/:id", h.DeleteLesson)
	lesson.GET("/list", h.ListLessons)

	// enrollment API

	enrollment := router.Group("/enrollment")
	enrollment.POST("/create", h.CreateEnrollment)
	enrollment.GET("read", h.ReadEnrollment)
	enrollment.PUT("update", h.UpdateEnrollment)
	enrollment.DELETE("/delete", h.DeleteEnrollment)

	router.Run("localhost:8080")
}
