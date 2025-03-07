package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/newWebBoxes/server/model"
	"github.com/sabrodigan/newWebBoxes/server/service"
	"github.com/sabrodigan/newWebBoxes/server/utils"
)

type UserController struct {
	userService     service.IUserService
	responseService utils.ResponseService
}

// CreateUser handles the creation of a new user from the provided JSON request and sends a success response or an error.
func (uc *UserController) CreateUser(ctx *gin.Context) {

	var dto model.UserCreateDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.Error(fmt.Errorf("400 :: %s :: %s :: %v", "invalid request", "userController.CreateUser", err))
		return
	}
	data, err := uc.userService.CreateUser(dto, nil)
	if err != nil {
		ctx.Error(fmt.Errorf("400 :: %s :: %s :: %v", "invalid request", "userController.CreateUser", err))
		return
	}

	uc.responseService.Success(ctx, 201, data, "successfully saved")

}
func GetUserController(userService service.IUserService, responseService utils.ResponseService) *UserController {
	return &UserController{
		userService:     userService,
		responseService: responseService,
	}
}
