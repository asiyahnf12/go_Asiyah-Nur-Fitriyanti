package controller

import (
	"mini_project/model"
	"mini_project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetOrdersController(c echo.Context) error {
	orders, e := usecase.GetListOrders()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"orders": orders,
	})
}

func GetOrderController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	order, err := usecase.GetOrder(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get order",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"orders": order,
	})
}

func CreateOrderController(c echo.Context) error {
	order := model.Order{}
	c.Bind(&order)

	if err := usecase.CreateOrder(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create order",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new order",
		"order":   order,
	})
}

func DeleteOrderController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteOrder(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete order",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf order tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete order",
	})
}

func UpdateOrderController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	order := model.Order{}
	c.Bind(&order)
	order.ID = uint(id)
	if err := usecase.UpdateOrder(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update order",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf order tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update order",
	})
}
