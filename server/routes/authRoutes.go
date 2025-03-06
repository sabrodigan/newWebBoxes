package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/newWebBoxes/server/controller"
	repo "github.com/sabrodigan/newWebBoxes/server/repository"
	"github.com/sabrodigan/newWebBoxes/server/service"
	"github.com/sabrodigan/newWebBoxes/server/utils"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	userService := service.GetUserService(*repository)              // Ensure correct arguments
	authService := service.GetAuthService(*repository, userService) // Ensure correct arguments
	responseService := utils.GetResponseService()
	authController := controller.GetAuthController(authService, *responseService) // Ensure correct type

	router.POST(
		"/login",
		authController.Login,
	)
}
