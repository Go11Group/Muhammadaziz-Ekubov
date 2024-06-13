package handLer

import (
	"add/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateCourse(c *gin.Context) {
	var course model.Course

	err := c.BindJSON(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.Course.CreateCourse(&course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusCreated, "Course created successfully")
}

func (h *Handler) ReadCoursesByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "ID not found in request")
	}
	course, err := h.Course.ReadCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"course": course})
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	var course model.Course

	err := c.BindJSON(&course)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err = h.Course.UpdateCourse(&course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, "Course updated successfully")
}

func (h *Handler) DeleteCourseByID(c *gin.Context) {
	id := c.Param("id")

	err := h.Course.DeleteCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, "Course deleted successfully")
}

func (h *Handler) listCourses(c *gin.Context) {
	courses := model.Courses{}
	courses.Title = c.Query("title")
	courses.Description = c.Query("description")

	c.JSON(http.StatusOK, gin.H{"courses": courses})
}
