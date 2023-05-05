package routes

import (
	"go-latihan/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	// Add other routes like GetUserController, CreateUserController, etc.

	return e
}
