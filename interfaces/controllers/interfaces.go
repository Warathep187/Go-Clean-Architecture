package controllers

import (
	"github.com/gin-gonic/gin"
)

type BlogController interface {
	CreateNewBlog(c *gin.Context)
	GetAllBlogs(c *gin.Context)
}

type UserController interface {
	RegisterUser(c *gin.Context)
}
