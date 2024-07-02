package CRUD

import (
	"add/model"
	"add/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	_, err = db.Exec("insert into Users(id ,progress,username, email, password) values ($1, $2, $3, $4, $5)", user.ID, user.Progress, user.Username, user.Email, user.Password)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusCreated, user)
}

func (h *Handler) GetUser(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var user model.User

	_, err = db.Exec("select * from Users where id=$1", c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var user model.User

	query := `update Users Set where progress=$1 , username=$2, email=$3, password=$4 where id=$5`

	_, err = db.Exec(query, user.Progress, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	db, err := postgres.ConnectDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var user model.User

	query := `update Users Set where progress=$1 , username=$2, email=$3, password=$4 where id=$5`

	_, err = db.Exec(query, user.Progress, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, user)
}
