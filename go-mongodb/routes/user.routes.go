package routes

import (
	"github.com/KemalBekir/Go-Tutorials/go-mongodb/controllers"
	"github.com/KemalBekir/Go-Tutorials/go-mongodb/middleware"
	"github.com/KemalBekir/Go-Tutorials/go-mongodb/services"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("user")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}
