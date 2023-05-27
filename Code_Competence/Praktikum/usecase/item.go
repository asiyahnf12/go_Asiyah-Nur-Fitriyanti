package usecase

import (
	"code_competence/model"
	"code_competence/repository/database"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type ItemUsecase interface {
	CreateItem(item *model.Item) error
	GetItem(id uint) (item model.Item, err error)
	GetListItems() (items []model.Item, err error)
	UpdateItem(item *model.Item) (err error)
	DeleteItem(id uint) (err error)
}

func CreateItem(item *model.Item) error {

	if item.Name == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Description == "" {
		return errors.New("item description not found")
	}

	err := database.CreateItem(item)
	if err != nil {
		return err
	}

	return nil
}

func GetItem(id uint) (item model.Item, err error) {
	item, err = database.GetItem(id)
	if err != nil {
		fmt.Println("GetItem: Error getting item from database")
		return
	}
	return
}

func GetListItems() (items []model.Item, err error) {
	items, err = database.GetItems()
	if err != nil {
		fmt.Println("GetListItems: Error getting items from database")
		return
	}
	return
}

func UpdateItem(item *model.Item) (err error) {
	err = database.UpdateItem(item)
	if err != nil {
		fmt.Println("UpdateItem : Error updating item, err: ", err)
		return
	}

	return
}

func DeleteItem(id uint) (err error) {
	item := model.Item{}
	item.ID = uuid.UUID{byte(id)}
	err = database.DeleteItem(&item)
	if err != nil {
		fmt.Println("DeleteItem : error deleting item, err: ", err)
		return
	}

	return
}
