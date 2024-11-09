package repositories

import (
	"context"
	"go-clean-arch/entities"
	"go-clean-arch/utils"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type blogRepository struct{}

func NewBlogRepository() BlogRepository {
	return &blogRepository{}
}

func (r *blogRepository) CreateBlog(in *entities.CreateBlogData) error {
	document := bson.D{
		{Key: "_id", Value: utils.GenerateRandomID()},
		{Key: "title", Value: in.Title},
		{Key: "content", Value: in.Content},
	}
	if _, err := mgm.Coll(&entities.Blog{}).InsertOne(context.Background(), &document); err != nil {
		return err
	}

	return nil
}

func (r *blogRepository) GetBlogs() ([]*entities.Blog, error) {
	var blogs []*entities.Blog

	cursor, err := mgm.Coll(&entities.Blog{}).Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &blogs); err != nil {
		return nil, err
	}
	if blogs == nil {
		return []*entities.Blog{}, nil
	}

	return blogs, nil
}
