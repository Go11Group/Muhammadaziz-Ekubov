package model

import (
	"time"
)

type Enrollment struct {
	EnrollmentID   string    `json:"enrollment_id"`
	UserID         string    `json:"user_id"`
	CourseID       string    `json:"course_id"`
	EnrollmentDate string    `json:"enrollment_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      int       `json:"deleted_at,omitempty"`
}
