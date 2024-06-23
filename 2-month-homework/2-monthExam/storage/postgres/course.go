package postgres

import (
	"add/model"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type CourseRepo struct {
	Db *sql.DB
}

// NewCourseRepo initializes a new CourseRepo with a given database connection.
func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{Db: db}
}

// CreateCourse inserts a new course into the database with a generated unique ID.
func (cr *CourseRepo) CreateCourse(course *model.Course) error {
	// Generate a new unique ID for the course.
	course.CourseID = uuid.NewString()

	// Execute the insert query to add the course to the database.
	_, err := cr.Db.Exec("insert into Courses (course_id, title, description) VALUES ($1, $2, $3)",
		&course.CourseID, &course.Title, &course.Description)
	if err != nil {
		return err
	}
	return nil
}

// ReadCourseByID retrieves a course from the database by its unique ID.
func (cr *CourseRepo) ReadCourseByID(CourseID string) (*model.Course, error) {
	var course model.Course

	// Execute the select query to fetch the course details from the database.
	rows := cr.Db.QueryRow("select * from Courses where course_id=$1", CourseID)

	// Scan the result into the course struct.
	err := rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &course, nil
}

// UpdateCourse updates the title and description of a course in the database.
func (cr *CourseRepo) UpdateCourse(course *model.Course) error {
	// Generate a new unique ID for the course.
	course.CourseID = uuid.NewString()

	// Execute the update query to modify the course details in the database.
	_, err := cr.Db.Exec("update Courses set title = $1, description = $2 where course_id=$3", course.Title, course.Description, course.CourseID)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCourse marks a course as deleted by setting the deleted_at timestamp.
func (cr *CourseRepo) DeleteCourse(courseID string) error {
	// Get the current time and convert it to Unix timestamp.
	now := time.Now()
	d := now.Unix()

	// Execute the update query to set the deleted_at timestamp for the course.
	_, err := cr.Db.Exec("update Courses set deleted_at=$1 where course_id = $2", d, courseID)
	if err != nil {
		return err
	}

	return nil
}

// ReadAllCourses retrieves all courses from the database that match the given filters.
func (cr *CourseRepo) ReadAllCourses(f *model.CourseFilter) ([]model.Course, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)
	filter := ""

	// Base query to select courses that are not deleted.
	query := `select course_id, title, description
	 	from Courses where deleted_at = 0`

	// Apply filters based on the provided filter parameters.
	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title"
	}
	if len(f.Description) > 0 {
		params["description"] = f.Description
		filter += " and description = :description"
	}

	// Construct the final query by combining the base query, filters, and limit.
	query = query + filter + limit

	// Replace query parameters with actual values.
	query, arr = ReplaceQueryParams(query, params)

	// Execute the query to fetch the courses from the database.
	rows, err := cr.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Course

	// Iterate over the result set and populate the courses slice.
	for rows.Next() {
		var course model.Course
		rows.Scan(&course.CourseID, &course.Title, &course.Description)
		courses = append(courses, course)
	}

	// Check for any errors that occurred during iteration.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
