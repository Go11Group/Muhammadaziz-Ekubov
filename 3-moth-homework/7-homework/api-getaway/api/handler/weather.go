package handler

import (
	"Muhammadaziz-Ekubov/3-moth-homework/7-homework/second_service/genproto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Get(ctx *gin.Context) {
	req := &genproto.Void{}

	resp, err := h.WhClient.(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

