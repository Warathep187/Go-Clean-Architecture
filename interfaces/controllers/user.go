package controllers

import (
	"go-clean-arch/models"
	"go-clean-arch/usecases"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) UserController {
	return &userController{userUsecase: userUsecase}
}

func (ctr *userController) RegisterUser(c *fiber.Ctx) error {
	var userData *models.CreateUserDto
	err := c.BodyParser(&userData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	httpStatus, err := ctr.userUsecase.RegisterUser(userData)
	if err != nil {
		return c.Status(httpStatus).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(httpStatus).JSON(fiber.Map{"message": "User registered successfully"})
}
