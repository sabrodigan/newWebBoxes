package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/newWebBoxes/server/controller"
	repo "github.com/sabrodigan/newWebBoxes/server/repository"
	"github.com/sabrodigan/newWebBoxes/server/service"
	"github.com/sabrodigan/newWebBoxes/server/utils"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	userService := service.GetUserService(*repository)
	responseService := utils.GetResponseService()
	userController := controller.GetUserController(userService, *responseService)

	router.POST(
		"/add",
		userController.CreateUser,
	)

}
