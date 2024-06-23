package handLer

import (
	"add/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.User.CreateUser(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusCreated, user)

}

func (h *Handler) ReadUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found in request")
	}
	user, err := h.User.ReadUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user model.User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.User.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, "Successfully updated user")
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found in request")
	}

	err := h.User.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(200, "Successfully deleted user")
}

func (h *Handler) ListUser(c *gin.Context) {
	user := model.Users{}
	user.UserID = c.Param("id")
	user.Name = c.Query("name")
	user.Email = c.Query("email")
	user.Password = c.Query("password")

	c.JSON(http.StatusOK, user)
}
