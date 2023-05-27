package usecase

import (
	"code_competence/model"
	"code_competence/model/payload"
	"code_competence/repository/database"

	"github.com/google/uuid"
)

type ItemUseCase interface {
	CreateItem(req payload.CreateItemRequest) (*model.Item, error)
	GetItemByName(name string) ([]model.Item, error)
	GetAllItems() ([]model.Item, error)
	GetItemByID(id string) (*model.Item, error)
	UpdateItemByID(id string, item model.Item) (*model.Item, error)
	DeleteItemByID(id string) error
}

type itemUseCase struct {
	itemRepository database.ItemRepository
}

func NewItemUseCase(itemRepository database.ItemRepository) *itemUseCase {
	return &itemUseCase{itemRepository: itemRepository}
}

func (i *itemUseCase) CreateItem(req payload.CreateItemRequest) (*model.Item, error) {
	id, _ := uuid.NewRandom()
	item := &model.Item{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	item, err := i.itemRepository.CreateItem(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemUseCase) GetItemByName(name string) ([]model.Item, error) {
	item, err := i.itemRepository.GetItemByName(name)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemUseCase) GetAllItems() ([]model.Item, error) {
	items, err := i.itemRepository.GetAllItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemUseCase) GetItemByID(id string) (*model.Item, error) {
	item, err := i.itemRepository.GetItemByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemUseCase) UpdateItemByID(id string, item model.Item) (*model.Item, error) {
	items, err := i.itemRepository.GetItemByID(id)
	if err != nil {
		return nil, err
	}
	items.Name = item.Name
	items.Description = item.Description
	items.Price = item.Price
	items.Stock = item.Stock
	items.CategoryID = item.CategoryID
	newItem, err := i.itemRepository.UpdateItem(items)
	if err != nil {
		return nil, err
	}
	return newItem, nil

}

func (i *itemUseCase) DeleteItemByID(id string) error {
	items, err := i.itemRepository.GetItemByID(id)
	if err != nil {
		return err
	}

	err = i.itemRepository.DeleteItem(items)
	if err != nil {
		return err
	}
	return nil
}
