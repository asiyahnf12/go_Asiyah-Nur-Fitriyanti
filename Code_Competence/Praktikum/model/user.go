package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `gorm:"-"`
}
