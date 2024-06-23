package handLer

import (
	"add/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateLesson(c *gin.Context) {
	var lesson model.Lesson

	err := c.BindJSON(&lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = h.Lesson.CreateLesson(&lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"lesson": lesson})
}

func (h *Handler) ReadLessonsByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}

	lesson, err := h.Lesson.ReadLessonByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"lesson": lesson})

}

func (h *Handler) UpdateLesson(c *gin.Context) {
	var lesson model.Lesson

	err := c.BindJSON(&lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.Lesson.UpdateLesson(&lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"lesson": lesson})
}

func (h *Handler) DeleteLesson(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}
	err := h.Lesson.DeleteLesson(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted lesson"})
}

func (h *Handler) ListLessons(c *gin.Context) {
	lesson := model.Lessons{}
	lesson.Title = c.Query("title")
	lesson.Content = c.Query("content")

	c.JSON(http.StatusOK, gin.H{"message": "We get all"})
}
