package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreateMenu(menu *model.Menu) error {
	if err := config.DB.Create(menu).Error; err != nil {
		return err
	}
	return nil
}

func GetListMenus() (menus []model.Menu, err error) {
	if err = config.DB.Find(&menus).Error; err != nil {
		return
	}
	return
}

func GetMenu(id uint) (menu model.Menu, err error) {
	menu.ID = id
	if err = config.DB.First(&menu).Error; err != nil {
		return
	}
	return
}

func UpdateMenu(menu *model.Menu) error {
	if err := config.DB.Updates(menu).Error; err != nil {
		return err
	}
	return nil
}

func DeleteMenu(menu *model.Menu) error {
	if err := config.DB.Delete(menu).Error; err != nil {
		return err
	}
	return nil
}
