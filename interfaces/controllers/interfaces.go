package controllers

import "github.com/gofiber/fiber/v2"

type BlogController interface {
	CreateNewBlog(c *fiber.Ctx) error
	GetAllBlogs(c *fiber.Ctx) error
}

type UserController interface {
	RegisterUser(c *fiber.Ctx) error
}
