package route

import (
	"code_competence/controller"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	user := e.Group("/users")
	user.GET("", controller.GetUserController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.PUT("/:id", controller.UpdateUserController)
	user.DELETE("/:id", controller.DeleteUserController)

	category := e.Group("/categorys")
	category.GET("", controller.GetCategoryController)
	category.GET("/:id", controller.GetCategoryController)
	category.POST("", controller.CreateCategoryController)
	category.PUT("/:id", controller.UpdateCategoryController)
	category.DELETE("/:id", controller.DeleteCategoryController)

	item := e.Group("/items")
	item.GET("", controller.GetItemController)
	item.GET("/:id", controller.GetItemController)
	item.POST("", controller.CreateItemController)
	item.PUT("/:id", controller.UpdateItemController)
	item.DELETE("/:id", controller.DeleteItemController)

}
