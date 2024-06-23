package model

import (
	"time"
)

type Course struct {
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   int       `json:"deleted_at"`
}

type Courses struct {
	Title       string
	Description string
}
