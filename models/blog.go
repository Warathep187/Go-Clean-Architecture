package models

type CreateBlogDTO struct {
	UserID  uint   `json:"userId" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
