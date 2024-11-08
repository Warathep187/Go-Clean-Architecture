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

func TestRegisterUserFailedGetUserFailed(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("GetUserByUsername", mock.Anything).Return(nil, errors.New("Something Failed"))

	userUsecase := usecases.NewUserUsecase(
		userRepositoryMock,
	)

	httpStatus, err := userUsecase.RegisterUser(&models.CreateUserDto{
		Username: "test",
		Password: "test",
	})
	assert.Equal(t, httpStatus, constants.StatusInternalServerError, "Status should be 500")
	assert.Equal(t, err, errors.New("Something Failed"))

	userRepositoryMock.AssertExpectations(t)
}

func TestRegisterUserFailedUsernameAlreadyExists(t *testing.T) {
	username := "test"
	password := "test"

	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("GetUserByUsername", username).Return(&entities.User{ID: 1, Username: username, Password: password}, nil)

	userUsecasse := usecases.NewUserUsecase(
		userRepositoryMock,
	)

	httpStatus, err := userUsecasse.RegisterUser(&models.CreateUserDto{
		Username: username,
		Password: password,
	})
	assert.Equal(t, httpStatus, constants.StatusConflict, "Status should be 409")
	assert.Equal(t, err, errors.New("Username already exists"))

	userRepositoryMock.AssertExpectations(t)
}

func TestRegisterUserSuccess(t *testing.T) {
	userRepositoryMock := new(mocks.UserRepository)
	userRepositoryMock.On("GetUserByUsername", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
	userRepositoryMock.On("CreateUser", mock.Anything).Return(nil)

	userUsecase := usecases.NewUserUsecase(
		userRepositoryMock,
	)

	httpStatus, err := userUsecase.RegisterUser(&models.CreateUserDto{
		Username: "test",
		Password: "test",
	})
	assert.Equal(t, httpStatus, constants.StatusCreated, "Status should be 201")
	assert.Equal(t, err, nil)

	userRepositoryMock.AssertExpectations(t)
}
