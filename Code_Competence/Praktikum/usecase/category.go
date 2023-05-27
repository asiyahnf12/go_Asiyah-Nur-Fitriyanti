package usecase

import (
	"code_competence/model"
	"code_competence/model/payload"
	"code_competence/repository/database"
)

type CategoryUsecase interface {
	CreateCategory(req payload.CategoryRequest) (*payload.CategoryResponse, error)
	GetAllItemByCategoryID(id string) ([]model.Category, error)
}

type categoryUsecase struct {
	categorysRepository database.CategoryRepository
}

func NewCategoryUsecase(categorysRepository database.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categorysRepository: categorysRepository}
}

func (c *categoryUsecase) CreateCategory(req payload.CategoryRequest) (*payload.CategoryResponse, error) {
	newCategory := &model.Category{
		Name: req.Name,
	}
	category, err := c.categorysRepository.CreateCategory(newCategory)
	if err != nil {
		return nil, err
	}
	resp := payload.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return &resp, nil
}

func (c *categoryUsecase) GetAllItemByCategoryID(id string) ([]model.Category, error) {
	category, err := c.categorysRepository.GetAllItemByCategoryID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
