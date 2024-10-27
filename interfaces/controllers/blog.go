package controllers

import (
	"go-clean-arch/models"
	"go-clean-arch/usecases"

	"github.com/gofiber/fiber/v2"
)

type blogController struct {
	blogUsecase usecases.BlogUsecase
}

func NewBlogController(blogUsecase usecases.BlogUsecase) BlogController {
	return &blogController{
		blogUsecase: blogUsecase,
	}
}

func (ctr *blogController) CreateNewBlog(c *fiber.Ctx) error {
	var blogData *models.CreateBlogDTO
	err := c.BodyParser(&blogData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = ctr.blogUsecase.CreateBlog(blogData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Blog created successfully"})
}

func (ctr *blogController) GetAllBlogs(c *fiber.Ctx) error {
	blogs, err := ctr.blogUsecase.GetAllBlogs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(blogs)
}
