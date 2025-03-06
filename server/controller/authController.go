package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sabrodigan/newWebBoxes/server/dto"
	"github.com/sabrodigan/newWebBoxes/server/service"
	"github.com/sabrodigan/newWebBoxes/server/utils"
)

type AuthController struct {
	authService     service.IAuthService
	responseService utils.ResponseService
}

func (ac *AuthController) Login(ctx *gin.Context) {

	var loginDto dto.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "no user found", "authController.login", err))
		if err != nil {
			log.Fatalf("Error binding login dto: %v", err)
		}
		return
	}
	data, err := ac.authService.Login(loginDto, nil)
	if err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "no user found", "authController.login", err))
		return
	}

	ac.responseService.Success(ctx, 201, data, "successfully saved")
}

func GetAuthController(authService service.IAuthService, responseService utils.ResponseService) *AuthController {
	return &AuthController{
		authService:     authService,
		responseService: responseService,
	}
}
