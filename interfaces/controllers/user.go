package controllers

import (
	"go-clean-arch/constants"
	"go-clean-arch/models"
	"go-clean-arch/usecases"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) UserController {
	return &userController{userUsecase: userUsecase}
}

func (ctr *userController) RegisterUser(c *gin.Context) {
	var userData *models.CreateUserDto
	err := c.BindJSON(&userData)
	if err != nil {
		c.JSON(constants.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	httpStatus, err := ctr.userUsecase.RegisterUser(userData)
	if err != nil {
		c.JSON(httpStatus, gin.H{"message": err.Error()})
		return
	}

	c.JSON(httpStatus, gin.H{"message": "User registered successfully"})
}
