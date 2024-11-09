package middlewares

import (
	"go-clean-arch/constants"
	"go-clean-arch/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateBlogData(c *gin.Context) {
	var blogData *models.CreateBlogDTO
	err := c.BindJSON(&blogData)

	if err != nil {
		c.JSON(constants.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	validate := validator.New()
	err = validate.Struct(blogData)
	if err != nil {
		c.JSON(constants.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Set("blogData", blogData)
	c.Next()
}
