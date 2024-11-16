package repositories

import (
	"go-clean-arch/database"
	"go-clean-arch/entities"
)

type blogRepository struct {
	db database.Database
}

func NewBlogRepository(db database.Database) BlogRepository {
	return &blogRepository{db: db}
}

func (r *blogRepository) CreateBlog(in *entities.CreateBlogData) error {
	data := &entities.Blog{
		Title:   in.Title,
		Content: in.Content,
	}

	result := r.db.GetDb().Create(data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *blogRepository) GetBlogs() ([]*entities.Blog, error) {
	var blogs []*entities.Blog

	result := r.db.GetDb().Find(&blogs)

	if result.Error != nil {
		return nil, result.Error
	}

	return blogs, nil
}

// FOR TESTING
func (r *blogRepository) DeleteBlogs() error {
	result := r.db.GetDb().Exec("DELETE FROM blogs")

	if result.Error != nil {
		return result.Error
	}

	return nil
}
