package usecases

import (
	"go-clean-arch/entities"
	"go-clean-arch/models"
)

type BlogUsecase interface {
	CreateBlog(data *models.CreateBlogDTO) (int, error)
	GetAllBlogs() ([]*entities.Blog, int, error)
}

type UserUsecase interface {
	RegisterUser(data *models.CreateUserDto) (int, error)
}
