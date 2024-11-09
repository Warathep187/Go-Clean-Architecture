package entities

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:"-" json:"-"`
	ID               int    `bson:"_id" json:"id"`
	Username         string `bson:"username" json:"username"`
	Password         string `bson:"password" json:"password"`
}

type CreateUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
