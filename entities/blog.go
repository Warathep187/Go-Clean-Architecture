package entities

import (
	"github.com/kamva/mgm/v3"
)

type Blog struct {
	mgm.DefaultModel `bson:"-" json:"-"`
	ID               int    `bson:"_id" json:"id"`
	Title            string `bson:"title" json:"title"`
	Content          string `bson:"content" json:"content"`
}

type CreateBlogData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
