package repositories

import (
	"errors"
	"go-clean-arch/database"
	"go-clean-arch/entities"

	"gorm.io/gorm"
)

type userRepository struct {
	db database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(in *entities.CreateUserData) error {
	userData := &entities.User{
		Username: in.Username,
		Password: in.Password,
	}

	result := r.db.GetDb().Create(userData)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *userRepository) GetUserByID(id uint) (*entities.User, error) {
	var user *entities.User
	result := r.db.GetDb().First(&user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*entities.User, error) {
	var user *entities.User
	result := r.db.GetDb().Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return user, nil
}
