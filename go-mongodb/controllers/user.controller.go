package controllers

import (
	"net/http"

	"github.com/KemalBekir/Go-Tutorials/go-mongodb/models"
	"github.com/KemalBekir/Go-Tutorials/go-mongodb/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return NewUserController(userService)
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*models.DBResponse)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FIlteredResponse(currentUser)}})

}
