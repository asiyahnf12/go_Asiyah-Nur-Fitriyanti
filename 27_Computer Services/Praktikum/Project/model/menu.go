package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	MenuID      uint   `json:"menu_id"`
	NameMenu    string `json:"name_menu"`
	Price       uint   `json:"price"`
	Description string `json:"description"`
}
