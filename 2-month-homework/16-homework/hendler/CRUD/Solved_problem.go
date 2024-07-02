package CRUD

import (
	"add/model"
	"add/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSolvedProblems(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var solve model.SolvedProblem
	query := `insert into Solved_problems(user_id, problem_id)`
	_, err = db.Exec(query, solve.UserID, solve.ProblemID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Solved_problem created"})
}

func (h *Handler) GetSolvedProblems(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	query := `select * from Solved_problems`
	_, err = db.Exec(query)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Solved_problems found"})

}

func (h *Handler) UpdateSolvedProblems(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var solve model.SolvedProblem
	query := `update Solved_problems set problem_id = %1 where user_id = $2`
	_, err = db.Exec(query, solve.ProblemID, solve.UserID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Solved_problem updated"})
}

func (h *Handler) DeleteSolvedProblems(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var solve model.SolvedProblem
	query := `delete from Solved_problems where user_id = $1`
	_, err = db.Exec(query, solve.UserID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Solved_problem deleted"})
}
