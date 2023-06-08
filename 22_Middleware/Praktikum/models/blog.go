package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	UserId int    `json:"user_id"  form:"user_id"`
	Judul  string `json:"judul"    form:"judul"`
	Konten string `json:"konten"   form:"konten"`
	User   User
}
