package handLer

import (
	"add/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateEnrollment(c *gin.Context) {
	var enrollment model.Enrollment
	err := c.BindJSON(&enrollment)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	err = h.Enrollment.CreateEnrollment(&enrollment)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusCreated, enrollment)
}

func (h *Handler) ReadEnrollment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found in request")
	}
	enrollment, err := h.User.ReadUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"user": enrollment})
}

func (h *Handler) UpdateEnrollment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found in request")
	}
	var enrollment model.Enrollment
	err := c.BindJSON(&enrollment)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"id": id, "enrollment": enrollment})
}

func (h *Handler) DeleteEnrollment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found in request")
	}
	err := h.Enrollment.DeleteEnrollment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) ListEnrollment(c *gin.Context) {

}
