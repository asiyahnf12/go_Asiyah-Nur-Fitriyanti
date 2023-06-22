package route

import (
	"mini_project/controller"
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.CreateUserController)

	user := e.Group("/users")
	user.GET("", controller.GetUsersController)
	user.GET("/:id", controller.GetUserController)
	user.POST("", controller.CreateUserController)
	user.PUT("/:id", controller.UpdateUserController)
	user.DELETE("/:id", controller.DeleteUserController)

	menu := e.Group("/menus")
	menu.GET("", controller.GetMenuController)
	menu.GET("/:id", controller.GetMenuController)
	menu.POST("", controller.CreateMenuController)
	menu.PUT("/:id", controller.UpdateMenuController)
	menu.DELETE("/:id", controller.DeleteMenuController)

	order := e.Group("/orders")
	order.GET("", controller.GetOrderController)
	order.GET("/:id", controller.GetOrderController)
	order.POST("", controller.CreateOrderController)
	order.PUT("/:id", controller.UpdateOrderController)
	order.DELETE("/:id", controller.DeleteOrderController)

}
