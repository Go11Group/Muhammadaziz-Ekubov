package model

import (
	"time"
)

type Lesson struct {
	LessonID  string    `json:"lesson_id"`
	CourseID  string    `json:"course_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at,omitempty"`
}

type Lessons struct {
	Title   string
	Content string
}
