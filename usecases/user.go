package usecases

import (
	"errors"
	"go-clean-arch/constants"
	"go-clean-arch/entities"
	"go-clean-arch/models"
	databaseRepository "go-clean-arch/repositories/database"
)

type userUsecase struct {
	userRepo databaseRepository.UserRepository
}

func NewUserUsecase(userRepo databaseRepository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) RegisterUser(data *models.CreateUserDto) (int, error) {
	username := data.Username
	password := data.Password

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return constants.StatusInternalServerError, err
	}
	if user != nil {
		return constants.StatusConflict, errors.New("Username already exists")
	}

	if err = u.userRepo.CreateUser(&entities.CreateUserData{
		Username: username,
		Password: password,
	}); err != nil {
		return constants.StatusInternalServerError, err
	}

	return constants.StatusCreated, nil
}
