package usecases

import (
	"errors"
	"go-clean-arch/entities"
	"go-clean-arch/models"
	"go-clean-arch/repositories"

	"gorm.io/gorm"
)

type blogUsecase struct {
	blogRepo repositories.BlogRepository
	userRepo repositories.UserRepository
}

func NewBlogUsecase(blogRepo repositories.BlogRepository, userRepo repositories.UserRepository) BlogUsecase {
	return &blogUsecase{
		blogRepo: blogRepo,
		userRepo: userRepo,
	}
}

func (u *blogUsecase) CreateBlog(data *models.CreateBlogDTO) error {
	userID := data.UserID
	title := data.Title
	content := data.Content

	_, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("User not found. Cannot create blog.")
		}
		return err
	}

	return u.blogRepo.CreateBlog(&entities.CreateBlogData{
		Title:   title,
		Content: content,
	})
}

func (u *blogUsecase) GetAllBlogs() ([]*entities.Blog, error) {
	return u.blogRepo.GetBlogs()
}
