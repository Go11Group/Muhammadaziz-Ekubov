package model

import (
	"github.com/google/uuid"
)

type Problem struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Difficulty  string    `json:"difficulty"`
	IsSolved    bool      `json:"is_solved"`
}
