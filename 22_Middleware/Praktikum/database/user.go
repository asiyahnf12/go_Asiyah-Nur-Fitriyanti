package database

import (
	"go-latihan/config"
	"go-latihan/middlewares"
	"go-latihan/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if e := config.DB.Preload("Blogs").Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	var user models.User
	if e := config.DB.Preload("Blogs").Find(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	if e := config.DB.Create(user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUserById(id int) (interface{}, error) {
	var user models.User
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUser(user models.User, id any) (models.User, error) {
	err := config.DB.Where("id = ?", id).Updates(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User
	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}
