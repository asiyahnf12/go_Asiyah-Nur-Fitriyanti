package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primarykey" json:"id" `
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`
	Stock       uint      `json:"stock"`
	Price       uint      `json:"price"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"-"`
}
