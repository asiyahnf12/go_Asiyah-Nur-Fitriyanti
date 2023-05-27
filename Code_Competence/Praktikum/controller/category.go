package controller

import (
	"code_competence/model"
	"net/http"
	"strconv"

	"code_competence/usecase"

	"github.com/labstack/echo/v4"
)

func GetCategorycontroller(c echo.Context) error {
	categorys, e := usecase.GetListCategorys()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"categorys": categorys,
	})
}

func GetCategoryController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	category, err := usecase.GetCategory(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get category",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success",
		"categorys": category,
	})
}

func CreateCategoryController(c echo.Context) error {
	category := model.Category{}
	c.Bind(&category)

	if err := usecase.CreateCategory(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create category",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new category",
		"category": category,
	})
}

func DeleteCategoryController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteCategory(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete category",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf category tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete category",
	})
}

func UpdateCategoryController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	category := model.Category{}
	c.Bind(&category)
	category.ID = uint(id)
	if err := usecase.UpdateCategory(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update category",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf category tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update category",
	})
}
