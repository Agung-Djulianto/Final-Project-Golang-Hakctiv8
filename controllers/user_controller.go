package controllers

import (
	"Project-Akhir/models"
	"Project-Akhir/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(UserService services.UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var request models.UserRegisterRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.MyError{
			Err: err.Error(),
		})
		return
	}

	response, err := uc.UserService.Add(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.UserRegisterResponse{
		ID:        response.ID,
		UserName:  response.UserName,
		Email:     request.Email,
		Age:       request.Age,
		CreatedAt: response.CreatedAt,
	})

}

func (uc *UserController) Login(ctx *gin.Context) {
	var request models.UserLoginRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.MyError{
			Err: err.Error(),
		})
		return
	}

	response, err := uc.UserService.Login(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
