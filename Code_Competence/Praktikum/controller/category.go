package controller

import (
	"code_competence/middleware"
	"code_competence/model/payload"
	"code_competence/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController interface {
	CreateCategoryController(c echo.Context) error
	GetAllByCategoryIDController(c echo.Context) error
}

type categoryController struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryController(
	categoryUsecase usecase.CategoryUsecase,
) *categoryController {
	return &categoryController{
		categoryUsecase: categoryUsecase,
	}
}

func (cc *categoryController) CreateCategoryController(c echo.Context) error {
	_, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	req := payload.CategoryRequest{}
	c.Bind(&req)

	category, err := cc.categoryUsecase.CreateCategory(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Create Category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success Create Category",
		"category": category,
	})
}

func (cc *categoryController) GetAllByCategoryIDController(c echo.Context) error {
	_, err := middleware.ExtractTokenUser(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}
	id := c.Param("category_id")
	categories, err := cc.categoryUsecase.GetAllItemByCategoryID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to Get All Category",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success Get All Item By Category ID",
		"category": categories,
	})
}
