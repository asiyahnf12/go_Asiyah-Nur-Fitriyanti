package controller

import (
	"mini_project/model"
	"net/http"
	"strconv"

	"mini_project/usecase"

	"github.com/labstack/echo/v4"
)

func GetMenucontroller(c echo.Context) error {
	menus, e := usecase.GetListMenus()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"menus":  menus,
	})
}

func GetMenuController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	menu, err := usecase.GetMenu(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get menu",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"menus":  menu,
	})
}

func CreateMenuController(c echo.Context) error {
	menu := model.Menu{}
	c.Bind(&menu)

	if err := usecase.CreateMenu(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create menu",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new menu",
		"menu":    menu,
	})
}

func DeleteMenuController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteMenu(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete menu",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf menu tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete menu",
	})
}

func UpdateMenuController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	menu := model.Menu{}
	c.Bind(&menu)
	menu.ID = uint(id)
	if err := usecase.UpdateMenu(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update menu",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf menu tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update menu",
	})
}
