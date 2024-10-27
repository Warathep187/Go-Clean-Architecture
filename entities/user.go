package entities

type User struct {
	ID       uint   `gorm:"primaryKey,autoIncrement" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
