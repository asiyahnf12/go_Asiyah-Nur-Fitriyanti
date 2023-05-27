package controller

import (
	"code_competence/middleware"
	"code_competence/model"
	"code_competence/model/payload"
	"code_competence/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemController interface {
	CreateItemController(c echo.Context) error
	GetItemByItemNameController(c echo.Context) error
	GetAllItemsController(c echo.Context) error
	GetItemByIDController(c echo.Context) error
	UpdateItemByIDController(c echo.Context) error
	DeleteItemByIDController(c echo.Context) error
}

type itemController struct {
	itemsUsecase usecase.ItemUseCase
}

func NewItemController(itemsUsecase usecase.ItemUseCase) *itemController {
	return &itemController{itemsUsecase: itemsUsecase}
}

func (i *itemController) CreateItemController(c echo.Context) error {
	_, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}

	req := payload.CreateItemRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	item, err := i.itemsUsecase.CreateItem(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Create Item",
		"item":    item,
	})
}

func (i *itemController) GetItemByItemNameController(c echo.Context) error {
	_, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	name := c.QueryParam("name")

	item, err := i.itemsUsecase.GetItemByName(name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Get Item",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Success Get Items By Name",
		"item":    item,
	})
}

func (i *itemController) GetAllItemsController(c echo.Context) error {
	_, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	items, err := i.itemsUsecase.GetAllItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all items",
		"items":   items,
	})
}

func (i *itemController) GetItemByIDController(c echo.Context) error {
	_, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := c.Param("id")

	item, err := i.itemsUsecase.GetItemByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Get Item By ID",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Item By ID",
		"Item":    item,
	})
}

func (i *itemController) UpdateItemByIDController(c echo.Context) error {
	_, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := c.Param("id")
	req := model.Item{}
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	item, err := i.itemsUsecase.UpdateItemByID(id, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Update Item By ID",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Item By ID",
		"Item":    item,
	})
}

func (i *itemController) DeleteItemByIDController(c echo.Context) error {
	_, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}

	id := c.Param("id")

	err = i.itemsUsecase.DeleteItemByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "Item Not Found ",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Item",
	})
}
