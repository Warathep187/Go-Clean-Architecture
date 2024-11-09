package controllers

import (
	"go-clean-arch/models"
	"go-clean-arch/usecases"

	"github.com/gin-gonic/gin"
)

type blogController struct {
	blogUsecase usecases.BlogUsecase
}

func NewBlogController(blogUsecase usecases.BlogUsecase) BlogController {
	return &blogController{
		blogUsecase: blogUsecase,
	}
}

func (ctr *blogController) CreateNewBlog(c *gin.Context) {
	blogData := c.MustGet("blogData").(*models.CreateBlogDTO)

	httpStatus, err := ctr.blogUsecase.CreateBlog(blogData)
	if err != nil {
		c.JSON(httpStatus, gin.H{"message": err.Error()})
		return
	}

	c.JSON(httpStatus, gin.H{"message": "Blog created successfully"})
}

func (ctr *blogController) GetAllBlogs(c *gin.Context) {
	blogs, httpStatus, err := ctr.blogUsecase.GetAllBlogs()
	if err != nil {
		c.JSON(httpStatus, gin.H{"message": err.Error()})
		return
	}
	c.JSON(httpStatus, blogs)
}
