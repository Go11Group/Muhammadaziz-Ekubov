package CRUD

import (
	"add/model"
	"add/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProblem(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	var problem model.Problem

	query := `INSERT INTO Problems VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(query, problem.ID, problem.Description, problem.Difficulty, problem.IsSolved)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, problem)
}

func (h *Handler) GetProblem(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var problem model.Problem

	query := `SELECT * FROM Problems WHERE id = $1`
	_, err = db.Query(query, problem.ID)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, problem)
}

func (h *Handler) UpdateProblem(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var problem model.Problem
	query := `update Problems set description = $1, difficulty= $2 , IsSolved=$3, where id = $4`
	_, err = db.Exec(query, problem.Description, problem.Difficulty, problem.IsSolved, problem.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, problem)
}

func (h *Handler) DeleteProblem(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var problem model.Problem
	query := `DELETE FROM Problems WHERE id = $1`
	_, err = db.Exec(query, problem.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, problem)
}
