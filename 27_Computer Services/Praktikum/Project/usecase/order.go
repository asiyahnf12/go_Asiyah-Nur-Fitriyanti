package usecase

import (
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

type OrderUsecase interface {
	CreateOrder(order *model.Order) error
	GetOrder(id uint) (order model.Order, err error)
	GetListOrders() (orders []model.Order, err error)
	UpdateOrder(order *model.Order) (err error)
	DeleteOrder(id uint) (err error)
}

func CreateOrder(order *model.Order) error {
	err := database.CreateOrder(order)
	if err != nil {
		return err
	}

	return nil
}

func GetOrder(id uint) (order model.Order, err error) {
	order, err = database.GetOrder(id)
	if err != nil {
		fmt.Println("GetOrder: Error getting order from database")
		return
	}
	return
}

func GetListOrders() (orders []model.Order, err error) {
	orders, err = database.GetOrders()
	if err != nil {
		fmt.Println("GetListOrders: Error getting orders from database")
		return
	}
	return
}

func UpdateOrder(order *model.Order) (err error) {
	err = database.UpdateOrder(order)
	if err != nil {
		fmt.Println("UpdateOrder : Error updating order, err: ", err)
		return
	}

	return
}

func DeleteOrder(id uint) (err error) {
	order := model.Order{}
	order.ID = id
	err = database.DeleteOrder(&order)
	if err != nil {
		fmt.Println("DeleteOrder : error deleting order, err: ", err)
		return
	}

	return
}
