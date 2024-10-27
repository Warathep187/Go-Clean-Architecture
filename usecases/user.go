package usecases

import (
	"errors"
	"go-clean-arch/entities"
	"go-clean-arch/models"
	"go-clean-arch/repositories"

	"gorm.io/gorm"
)

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) RegisterUser(data *models.CreateUserDto) error {
	username := data.Username
	password := data.Password

	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if user != nil {
		return errors.New("Username already exists")
	}

	if err = u.userRepo.CreateUser(&entities.CreateUserData{
		Username: username,
		Password: password,
	}); err != nil {
		return err
	}

	return nil
}
