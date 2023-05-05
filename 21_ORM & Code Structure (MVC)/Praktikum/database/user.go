package database

import (
	"user/config"
	"user/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

// Add other database functions like GetUser, CreateUser, etc.
