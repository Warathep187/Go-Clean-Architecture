package middlewares

import (
	"go-clean-arch/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBlogData(c *fiber.Ctx) error {
	var blogData *models.CreateBlogDTO
	err := c.BodyParser(&blogData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	validate := validator.New()
	err = validate.Struct(blogData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Next()
}
