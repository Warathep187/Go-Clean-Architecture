package usecases

import (
	"errors"
	"go-clean-arch/constants"
	"go-clean-arch/entities"
	"go-clean-arch/models"
	databaseRepository "go-clean-arch/repositories/database"
)

type blogUsecase struct {
	blogRepo databaseRepository.BlogRepository
	userRepo databaseRepository.UserRepository
}

func NewBlogUsecase(blogRepo databaseRepository.BlogRepository, userRepo databaseRepository.UserRepository) BlogUsecase {
	return &blogUsecase{
		blogRepo: blogRepo,
		userRepo: userRepo,
	}
}

func (u *blogUsecase) CreateBlog(data *models.CreateBlogDTO) (int, error) {
	userID := data.UserID
	title := data.Title
	content := data.Content

	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return constants.StatusInternalServerError, err
	}
	if user == nil {
		return constants.StatusNotFound, errors.New("User not found. Cannot create blog.")
	}

	if err = u.blogRepo.CreateBlog(&entities.CreateBlogData{
		Title:   title,
		Content: content,
	}); err != nil {
		return constants.StatusInternalServerError, err
	}

	return constants.StatusCreated, nil
}

func (u *blogUsecase) GetAllBlogs() ([]*entities.Blog, int, error) {
	blogs, err := u.blogRepo.GetBlogs()
	if err != nil {
		return nil, constants.StatusInternalServerError, err
	}
	return blogs, constants.StatusOK, nil
}
