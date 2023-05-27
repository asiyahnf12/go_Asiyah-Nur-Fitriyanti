package database

import (
	"code_competence/config"
	"code_competence/model"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item *model.Item) (*model.Item, error)
	GetItemByName(name string) ([]model.Item, error)
	GetAllItems() ([]model.Item, error)
	GetItemByID(id string) (*model.Item, error)
	UpdateItem(item *model.Item) (*model.Item, error)
	DeleteItem(item *model.Item) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

func (i *itemRepository) CreateItem(item *model.Item) (*model.Item, error) {
	err := config.DB.Create(item).Error
	if err != nil {
		return nil, err
	}
	err = config.DB.Preload("Category").First(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemRepository) GetItemByName(name string) ([]model.Item, error) {
	var items []model.Item

	err := config.DB.Where("name LIKE ? ", "%"+name+"%").Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemRepository) GetAllItems() ([]model.Item, error) {
	var items []model.Item
	err := config.DB.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemRepository) GetItemByID(id string) (*model.Item, error) {
	var item model.Item
	err := config.DB.Where("id = ?", id).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (i *itemRepository) UpdateItem(item *model.Item) (*model.Item, error) {
	err := config.DB.Save(item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (i *itemRepository) DeleteItem(item *model.Item) error {
	err := config.DB.Delete(&item).Error
	if err != nil {
		return err
	}
	return nil
}
