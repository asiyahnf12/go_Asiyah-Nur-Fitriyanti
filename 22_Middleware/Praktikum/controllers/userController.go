package controllers

import (
	"go-latihan/database"
	"go-latihan/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	UserId := GetIdFromParam(c)
	user, e := database.GetUserById(UserId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"users":   user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := &models.User{}

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdUser, err := database.CreateUser(user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    createdUser,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	UserId := GetIdFromParam(c)
	user, e := database.DeleteUserById(UserId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Delete user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	UserId := GetIdFromParam(c)
	updatedUser := models.User{}

	if e := c.Bind(updatedUser); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	user, e := database.UpdateUser(updatedUser, UserId)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

func GetIdFromParam(c echo.Context) int {
	UserId, _ := strconv.Atoi(c.Param("id"))
	return UserId
}
