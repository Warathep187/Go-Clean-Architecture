package controllers

import (
	"go-clean-arch/constants"
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
		return c.Status(constants.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	httpStatus, err := ctr.blogUsecase.CreateBlog(blogData)
	if err != nil {
		return c.Status(httpStatus).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(httpStatus).JSON(fiber.Map{
		"message": "Blog created successfully",
	})
}

func (ctr *blogController) GetAllBlogs(c *fiber.Ctx) error {
	blogs, httpStatus, err := ctr.blogUsecase.GetAllBlogs()
	if err != nil {
		return c.Status(httpStatus).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(httpStatus).JSON(blogs)
}
