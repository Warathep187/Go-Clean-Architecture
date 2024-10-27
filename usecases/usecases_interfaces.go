package usecases

import (
	"go-clean-arch/entities"
	"go-clean-arch/models"
)

type BlogUsecase interface {
	CreateBlog(data *models.CreateBlogDTO) error
	GetAllBlogs() ([]*entities.Blog, error)
}

type UserUsecase interface {
	RegisterUser(data *models.CreateUserDto) error
}
