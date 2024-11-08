package test_usecases

import (
	"errors"
	"go-clean-arch/constants"
	"go-clean-arch/entities"
	"go-clean-arch/models"
	"go-clean-arch/tests/mocks"
	"go-clean-arch/usecases"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateBlogFailedUserNotFound(t *testing.T) {
	blogRepositoryMock := new(mocks.BlogRepository)
	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("GetUserByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	blogUsecase := usecases.NewBlogUsecase(
		blogRepositoryMock,
		userRepositoryMock,
	)

	blogData := models.CreateBlogDTO{
		UserID:  1,
		Title:   "Test Blog",
		Content: "Test Content",
	}
	httpStatus, err := blogUsecase.CreateBlog(&blogData)
	assert.Equal(t, httpStatus, constants.StatusNotFound, "Status should be 404")
	assert.Equal(
		t,
		err,
		errors.New("User not found. Cannot create blog."),
		"Error should be 'User not found. Cannot create blog.'",
	)

	userRepositoryMock.AssertExpectations(t)
}

func TestCreateBlogFailedGetUserFailed(t *testing.T) {
	blogRepositoryMock := new(mocks.BlogRepository)
	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("GetUserByID", mock.Anything).Return(nil, errors.New("Something Failed"))

	blogUsecase := usecases.NewBlogUsecase(
		blogRepositoryMock,
		userRepositoryMock,
	)

	blogData := models.CreateBlogDTO{
		UserID:  1,
		Title:   "Test Blog",
		Content: "Test Content",
	}
	httpStatus, err := blogUsecase.CreateBlog(&blogData)
	assert.Equal(t, httpStatus, constants.StatusInternalServerError, "Status should be 500")
	assert.Equal(
		t,
		err,
		errors.New("Something Failed"),
		"Error should be 'Something Failed'",
	)

	userRepositoryMock.AssertExpectations(t)
}

func TestCreateBlogSuccess(t *testing.T) {
	blogRepositoryMock := new(mocks.BlogRepository)
	blogRepositoryMock.On("CreateBlog", mock.Anything).Return(nil)
	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("GetUserByID", mock.Anything).Return(&entities.User{ID: 1}, nil)

	blogUsecase := usecases.NewBlogUsecase(
		blogRepositoryMock,
		userRepositoryMock,
	)

	blogData := models.CreateBlogDTO{
		UserID:  1,
		Title:   "Test Blog",
		Content: "Test Content",
	}
	httpStatus, err := blogUsecase.CreateBlog(&blogData)
	assert.Equal(t, httpStatus, constants.StatusCreated, "Status should be 201")
	assert.Equal(t, err, nil, "Error should be nil")

	blogRepositoryMock.AssertExpectations(t)
	userRepositoryMock.AssertExpectations(t)
}

func TestGetAllBlogsFailed(t *testing.T) {
	blogRepositoryMock := new(mocks.BlogRepository)
	blogRepositoryMock.On("GetBlogs").Return(nil, errors.New("Something Failed"))
	userRepositoryMock := new(mocks.UserRepository)

	blogUsecases := usecases.NewBlogUsecase(
		blogRepositoryMock,
		userRepositoryMock,
	)

	_, httpStatus, err := blogUsecases.GetAllBlogs()
	assert.Equal(
		t,
		httpStatus,
		constants.StatusInternalServerError,
		"Status should be 500",
	)
	assert.Equal(
		t,
		err,
		errors.New("Something Failed"),
		"Error should be 'Something Failed'",
	)

	blogRepositoryMock.AssertExpectations(t)
}

func TestGetAllBlogSuccess(t *testing.T) {
	blogsMock := []*entities.Blog{}
	blogsMock = append(blogsMock, &entities.Blog{ID: 1, Title: "Test Blog 1", Content: "Test Content 1"})

	blogRepositoryMock := new(mocks.BlogRepository)
	blogRepositoryMock.On("GetBlogs").Return(blogsMock, nil)
	userRepositoryMock := new(mocks.UserRepository)

	blogUsecases := usecases.NewBlogUsecase(
		blogRepositoryMock,
		userRepositoryMock,
	)

	blogs, httpStatus, _ := blogUsecases.GetAllBlogs()
	assert.Equal(t, httpStatus, constants.StatusOK, "Status should be 200")
	assert.Equal(t, blogs, blogsMock, "Returned blogs should be equal")

	blogRepositoryMock.AssertExpectations(t)
}
