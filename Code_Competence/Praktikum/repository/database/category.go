package database

import (
	"code_competence/config"
	"code_competence/model"
)

func CreateCategory(category *model.Category) error {
	if err := config.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func GetCategorys() (categorys []model.Category, err error) {
	if err = config.DB.Find(&categorys).Error; err != nil {
		return
	}
	return
}

func GetCategory(id uint) (category model.Category, err error) {
	category.ID = id
	if err = config.DB.First(&category).Error; err != nil {
		return
	}
	return
}

func UpdateCategory(category *model.Category) error {
	if err := config.DB.Updates(category).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCategory(category *model.Category) error {
	if err := config.DB.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
