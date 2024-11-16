package repositories

import "go-clean-arch/entities"

type BlogRepository interface {
	CreateBlog(dto *entities.CreateBlogData) error
	GetBlogs() ([]*entities.Blog, error)
	// FOR TESTING
	DeleteBlogs() error
}

type UserRepository interface {
	CreateUser(dto *entities.CreateUserData) error
	GetUserByID(id uint) (*entities.User, error)
	GetUserByUsername(username string) (*entities.User, error)
	// FOR TESTING
	CreateUserWithID(id uint, dto *entities.CreateUserData) error
	DeleteUsers() error
}
