package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StatusOK(c *gin.Context, err error) {
	c.JSON(200, gin.H{
		"status":  http.StatusOK,
		"time":    time.Now(),
		"message": err,
		"success": true,
	})
	c.Header("Content-Type", "application/json")

}
func StatusCreated(c *gin.Context, err error) {
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"time":    time.Now(),
		"message": err,
		"success": true,
	})
	c.Header("Content-Type", "application/json")

}
func StatusInternalServerError(c *gin.Context, err error) {
	fmt.Println("salom")
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"time":    time.Now(),
		"message": err,
		"success": false,
	})
	c.Header("Content-Type", "application/json")

}
func StatusBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"time":    time.Now(),
		"message": err,
		"success": false,
	})
	c.Header("Content-Type", "application/json")

}
