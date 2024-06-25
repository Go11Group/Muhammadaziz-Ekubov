package handler

import (
	"Muhammadaziz-Ekubov/3-moth-homework/3-homework/metro/storage/postgres"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Db   *sql.DB
	User *postgres.UserRepository
}

func (h *Handler) Routes() {
	router := gin.Default()

	users := router.Group("/users")

	users.GET("/all", h.GetAllUsers)
}
