package usecase

import (
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

func CreateMenu(Menu *model.Menu) error {
	err := database.CreateMenu(Menu)
	if err != nil {
		return err
	}

	return nil
}

func GetMenu(id uint) (menu model.Menu, err error) {
	menu, err = database.GetMenu(id)
	if err != nil {
		fmt.Println("GetMenu: Error getting menu from database")
		return
	}
	return
}

func GetListMenus() (menus []model.Menu, err error) {
	menus, err = database.GetListMenus()
	if err != nil {
		fmt.Println("GetListMenus: Error getting menus from database")
		return
	}
	return
}

func UpdateMenu(menu *model.Menu) (err error) {
	err = database.UpdateMenu(menu)
	if err != nil {
		fmt.Println("UpdateMenu : Error updating menu, err: ", err)
		return
	}

	return
}

func DeleteMenu(id uint) (err error) {
	menu := model.Menu{}
	menu.ID = id
	err = database.DeleteMenu(&menu)
	if err != nil {
		fmt.Println("DeleteMenu : error deleting menu, err: ", err)
		return
	}

	return
}
