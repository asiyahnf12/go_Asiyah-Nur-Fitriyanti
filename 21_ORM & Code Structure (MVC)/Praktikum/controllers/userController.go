package controllers

import (
	"go-latihan/database"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error getting users",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// Add other controllers like GetUserController, CreateUserController, etc.
