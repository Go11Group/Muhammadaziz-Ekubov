package model

import "github.com/google/uuid"

type SolvedProblem struct {
	UserID    uuid.UUID `json:"user_id"`
	ProblemID uuid.UUID `json:"problem_id"`
}
