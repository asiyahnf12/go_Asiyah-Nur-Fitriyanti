package database

import (
	"code_competence/config"
	"code_competence/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(req *model.Category) (*model.Category, error)
	GetAllItemByCategoryID(id string) ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) CreateCategory(category *model.Category) (*model.Category, error) {
	err := config.DB.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *categoryRepository) GetAllItemByCategoryID(id string) ([]model.Category, error) {
	var category []model.Category
	err := config.DB.Preload("Item").Where("id = ?", id).Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
