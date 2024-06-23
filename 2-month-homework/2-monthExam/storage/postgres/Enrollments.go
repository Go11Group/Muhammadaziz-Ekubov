package postgres

import (
	"add/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type EnrollmentRepo struct {
	Db *sql.DB
}

type EnrollmentFilter struct {
	EnrollmentID   string
	UserID         string
	CourseID       string
	EnrollmentDate string
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{Db: db}
}

func (e *EnrollmentRepo) CreateEnrollment(enrollment *model.Enrollment) error {
	enrollmentDate, err := time.Parse("2006-01-02", enrollment.EnrollmentDate)
	if err != nil {
		return errors.New("failed to parse birthday")
	}
	enrollmentID := uuid.NewString()
	query := `insert into Enrollments (user_id, course_id, enrollment_date) values ($1, $2, $3)`
	_, err = e.Db.Exec(query, enrollmentID, enrollmentDate, enrollmentDate)

	if err != nil {
		return err
	}

	return nil
}

func (e *EnrollmentRepo) ReadEnrollmentsByID(EnrollmentID string) (*model.Enrollment, error) {
	var enrollment model.Enrollment
	row := e.Db.QueryRow("select * from Enrollments where enrollment_id = $1", EnrollmentID)

	err := row.Scan(
		&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &enrollment, nil
}

func (e *EnrollmentRepo) UpdateEnrollment(Enrollment *model.Enrollment) error {
	enrollmentDate, err := time.Parse("2006-01-02", Enrollment.EnrollmentDate)
	if err != nil {
		return errors.New("failed to parse birthday")
	}
	query := `update Enrollments set enrollment_date = $1 where enrollment_id = $2`
	_, err = e.Db.Exec(query, enrollmentDate, Enrollment.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EnrollmentRepo) DeleteEnrollment(EnrollmentID string) error {
	now := time.Now()
	deletedAt := now.Unix()

	query := `update Enrollments set deleted_at = $1 where enrollment_id = $2`

	_, err := e.Db.Exec(query, deletedAt, EnrollmentID)
	if err != nil {
		return err
	}
	return nil
}

func (e *EnrollmentRepo) GetAll(Enrollment *model.Enrollment) ([]model.Enrollment, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
	)
	f := EnrollmentFilter{}

	query := `select enrollment_id, user_id, course_id, enrollment_date
  from enrollments `
	filter := ` where true`

	if len(f.EnrollmentID) > 0 {
		params["enrollment_id"] = f.EnrollmentID
		filter += " and enrollment_id = :enrollment_id "
	}

	if len(f.UserID) > 0 {
		params["user_id"] = f.UserID
		filter += " and user_id = :user_id "
	}

	if len(f.CourseID) > 0 {
		params["course_id"] = f.CourseID
		filter += " and course_id = :course_id "
	}

	if len(f.EnrollmentDate) > 0 {
		params["enrollment_date"] = f.EnrollmentDate
		filter += " and enrollment_date = :enrollment_date "
	}

	query = query + filter

	query, arr = ReplaceQueryParams(query, params)
	rows, err := e.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.Enrollment
	for rows.Next() {
		var enrrolment model.Enrollment
		err = rows.Scan(&enrrolment.EnrollmentID, &enrrolment.UserID, &enrrolment.CourseID, &enrrolment.EnrollmentDate)
		if err != nil {
			return nil, err
		}

		enrollments = append(enrollments, enrrolment)
	}
	return enrollments, nil
}
