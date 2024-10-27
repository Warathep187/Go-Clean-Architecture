package entities

type Blog struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateBlogData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
