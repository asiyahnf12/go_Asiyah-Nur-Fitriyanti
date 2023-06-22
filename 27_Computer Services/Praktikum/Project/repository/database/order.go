package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreateOrder(order *model.Order) error {
	if err := config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrders() (orders []model.Order, err error) {
	if err = config.DB.Find(&orders).Error; err != nil {
		return
	}
	return
}

func GetOrder(id uint) (order model.Order, err error) {
	order.ID = id
	if err = config.DB.First(&order).Error; err != nil {
		return
	}
	return
}

func GetOrdersByUserId(userID uint) (order model.Order, err error) {
	order.UserID = userID
	if err = config.DB.Find(&order).Error; err != nil {
		return
	}
	return
}

func UpdateOrder(order *model.Order) error {
	if err := config.DB.Updates(order).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOrder(order *model.Order) error {
	if err := config.DB.Delete(order).Error; err != nil {
		return err
	}
	return nil
}
