package postgres

import (
	"add/model"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type LessonRepo struct {
	db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db}
}

func (l *LessonRepo) CreateLesson(lesson *model.Lesson) error {
	lesson.LessonID = uuid.NewString()

	err := l.db.QueryRow("insert into lessons(lesson_id, course_id, title, content) values ($1, $2, $3, %4)").Scan(lesson.LessonID, lesson.CourseID, lesson.Title, lesson.Content)
	if err != nil {
		return err
	}
	return nil

}

func (l *LessonRepo) ReadLessonByID(LessonID string) (*model.Lesson, error) {
	var lesson model.Lesson
	row := l.db.QueryRow("select * from lessons where lesson_id = $1", LessonID)

	err := row.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
		}
	}
	return &lesson, nil
}

func (l *LessonRepo) UpdateLesson(lesson *model.Lesson) error {
	query := `update lessons set title, content, updated_at where lesson_id = $1`
	_, err := l.db.Exec(query, &lesson.LessonID, &lesson.Title)
	if err != nil {
		return err
	}
	return nil
}

func (l *LessonRepo) DeleteLesson(LessonID string) error {
	now := time.Now()
	deletedAt := now.Unix()

	_, err := l.db.Exec("update lessons set deleted_at=$1 where lesson_id = $2", deletedAt, LessonID)
	if err != nil {
		return err
	}
	return nil
}

func (l *LessonRepo) ListLessons(f *model.LessonFilter) ([]model.Lessons, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
	)
	filter := ""

	query := `select title, content, created_at, updated_at
	 	from Lessons where deleted_at = 0 `

	if len(f.Title) > 50 {
		params["title"] = f.Title
		filter += " and title = :title"
	}
	if len(f.Content) > 250 {
		params["content"] = f.Content
		filter += " and content = :content"
	}

	query = query + filter + limit

	query, arr = ReplaceQueryParams(query, params)
	db := &sql.DB{}
	rows, err := db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var lessons []model.Lessons
	for rows.Next() {
		var lesson model.Lessons

		rows.Scan(&lesson.Title, &lesson.Content)

		lessons = append(lessons, lesson)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return lessons, nil
}
