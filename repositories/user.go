package repositories

import (
	"context"
	"errors"
	"go-clean-arch/entities"
	"go-clean-arch/utils"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(in *entities.CreateUserData) error {
	document := bson.D{
		{Key: "_id", Value: utils.GenerateRandomID()},
		{Key: "username", Value: in.Username},
		{Key: "password", Value: in.Password},
	}
	if _, err := mgm.Coll(&entities.User{}).InsertOne(context.Background(), &document); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByID(id uint) (*entities.User, error) {
	user := &entities.User{}

	if err := mgm.Coll(&entities.User{}).FindByID(id, user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*entities.User, error) {
	user := &entities.User{}

	query := bson.M{"username": username}
	if err := mgm.Coll(&entities.User{}).First(query, user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
