package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID     uint   `json:"order_id"`
	UserID      uint   `json:"user_id"`
	MenuID      uint   `json:"menu_id"`
	OrderStatus string `json:"status"`
}

type OrderDetail struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	OrderID   uint   `json:"order_id"`
	MenuID    uint   `json:"menu_id"`
	Quantity  uint   `json:"quantity"`
	CreatedAt string `json:"time"`
}
