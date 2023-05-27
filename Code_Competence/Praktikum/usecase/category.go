package usecase

import (
	"code_competence/model"
	"code_competence/repository/database"
	"fmt"
)

func CreateCategory(req *model.Category) error {
	category := model.Category{
		Name: req.Name,
	}
	err := database.CreateCategory(&category)
	if err != nil {
		return err
	}
	return nil
}

func GetCategory(id uint) (category model.Category, err error) {
	category, err = database.GetCategory(id)
	if err != nil {
		fmt.Println("GetCategory: Error getting category from database")
		return
	}
	return
}

func GetListCategorys() (categorys []model.Category, err error) {
	categorys, err = database.GetCategorys()
	if err != nil {
		fmt.Println("GetListCategorys: Error getting categorys from database")
		return
	}
	return
}

func UpdateCategory(category *model.Category) (err error) {
	err = database.UpdateCategory(category)
	if err != nil {
		fmt.Println("UpdateCategory : Error updating category, err: ", err)
		return
	}

	return
}

func DeleteCategory(id uint) (err error) {
	category := model.Category{}
	category.ID = id
	err = database.DeleteCategory(&category)
	if err != nil {
		fmt.Println("DeleteCategory : error deleting category, err: ", err)
		return
	}

	return
}
