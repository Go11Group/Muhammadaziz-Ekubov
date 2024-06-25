package handler

import (
	"Muhammadaziz-Ekubov/3-moth-homework/3-homework/metro/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllUsers(c *gin.Context) {
	var user model.Users
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

func (h *Handler) CreateUser(gn *gin.Context) {
	user := model.Users{}
	err := gn.BindJSON(&user)
	if err != nil {
		StatusBadRequest(gn, err)
		return
	}

	err = h.User.CreateUser(&user)
	if err != nil {
		StatusInternalServerError(gn, err)

		return
	}

	StatusCreated(gn, err)
}
func (h *Handler) UpdateUser(gn *gin.Context) {
	user := model.Users{}
	err := gn.BindJSON(&user)
	if err != nil {

		StatusBadRequest(gn, err)
		return

	}

	err = h.User.UpdateUser(&user)
	if err != nil {
		StatusInternalServerError(gn, err)
		return
	}
	StatusOK(gn, err)
}

func (h *Handler) DeleteUser(gn *gin.Context) {
	err := h.User.DeleteUser(gn.Param("id"))
	if err != nil {
		StatusInternalServerError(gn, err)
		return

	}
	StatusOK(gn, err)

}

func (h *Handler) GetUserById(gn *gin.Context) {
	users, err := h.User.GetByID(gn.Param("id"))
	if err != nil {
		StatusInternalServerError(gn, err)

	}
	gn.JSON(200, gin.H{
		"users": users,
	})
}

//func (h *Handler) GetUserCard(gn *gin.Context) {
//	userCards, err := h.UserRepo.GetUserCard(gn.Param("id"))
//	if err != nil {
//		fmt.Println("++++++", err)
//		InternalServerError(gn, err)
//
//	}
//	gn.JSON(200, gin.H{
//		"userCards": userCards,
//	})
//}
//func (h *Handler) GetUserCardDeposit(gn *gin.Context) {
//	userCardDeposits, err := h.UserRepo.GetUserDepositCard(gn.Param("id"))
//	if err != nil {
//		fmt.Println("++++++", err)
//		InternalServerError(gn, err)
//
//	}
//	gn.JSON(200, gin.H{
//		"userCardDeposits": userCardDeposits,
//	})
//}
//func (h *Handler) GetUserCardCredit(gn *gin.Context) {
//	userCardCredits, err := h.UserRepo.GetUserCreditCard(gn.Param("id"))
//	if err != nil {
//		fmt.Println("++++++", err)
//		InternalServerError(gn, err)
//
//	}
//	gn.JSON(200, gin.H{
//		"userCardCredits": userCardCredits,
//	})
//}
