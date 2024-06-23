package main

//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func (h *Handler) CreateUser(c *gin.Context) {
//	var user models.Users
//	err := c.BindJSON(&user)
//	if err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//	}
//	err = h.User.CreateUser(&user)
//	if err != nil {
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//	c.JSON(http.StatusCreated, user)
//}
//func (h *Handler) GetCourseByUsers(c *gin.Context) {
//	userId, course, err := h.User.GetCourseByUser(c.Param("id"))
//	if err != nil {
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//	c.JSON(http.StatusOK, user)
//
//}
//
//func (h *Handler) UpdateUser(c *gin.Context) {
//	var user models.Users
//	err := c.Bind(&user)
//	if err != nil {
//		c.String(http.StatusBadRequest, err.Error())
//	}
//	err = h.User.UpdateUser(&user)
//	if err != nil {
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//
//	c.String(200, "SUCCESS")
//
//}
//
//func (h *Handler) GetAllUsers(c *gin.Context) {
//	users, err := h.User.ReadAllUsers()
//	if err != nil {
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//	c.JSON(http.StatusOK, users)
//
//}
//
//func (h *Handler) GetUserById(c *gin.Context) {
//	id := c.Param("user_id")
//	if id == "" {
//		c.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
//	}
//
//	user, err := h.User.ReadUser(id)
//	if err != nil {
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//	c.JSON(http.StatusOK, user)
//
//}
//
//func (h *Handler) DeleteUsers(c *gin.Context) {
//	id := c.Param("user_id")
//	if id == "" {
//		c.String(http.StatusBadRequest, "ID kiritlishda hatolik mavjud")
//	}
//
//	err := h.User.DeleteUser(id)
//	if err != nil {
//		c.String(http.StatusInternalServerError, err.Error())
//	}
//	c.String(200, "SUCCESS")
//
//}
//
//func (h *Handler) GetFilterUsers(c *gin.Context) {
//	userFilter := models.Users{}
//	userFilter.ID = c.Param("user_id")
//	userFilter.Name = c.Param("name")
//	userFilter.Email = c.Param("email")
//	userFilter.Birthday = c.Param("birthday")
//	userFilter.Password = c.Param("password")
//
//	c.JSON(http.StatusOK, userFilter)
//}
