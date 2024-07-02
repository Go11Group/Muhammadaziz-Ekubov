package main

import (
	"add/hendler/CRUD"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	h := CRUD.Newhandler()

	//Problem routes
	r.POST("/createUser", h.CreateUser)
	r.GET("/problems/get/{id}", h.GetProblem)
	r.PUT("/problems/update/{id}", h.UpdateProblem)
	r.DELETE("/problems/delete/{id}", h.DeleteProblem)
	// Users routes
	r.POST("/users", h.CreateUser)
	r.GET("/users/get/{id}", h.GetUser)
	r.PUT("/users/update/{id}", h.UpdateUser)
	r.DELETE("/users/delete/{id}", h.DeleteUser)

	// SolvedProblems routes
	r.POST("/solved_problems", h.CreateSolvedProblems)
	r.GET("/solved_problems", h.GetSolvedProblems)
	r.GET("/solved_problems/update/{user_id}/{problem_id}", h.UpdateSolvedProblems)
	r.DELETE("/solved_problems/delete/{user_id}/{problem_id}", h.DeleteSolvedProblems)

	r.Run("localhost:8070")
}
