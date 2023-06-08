package route

import (
	"go-latihan/controllers"
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

	e.POST("/login", controllers.LoginUserController)
	e.POST("/register", controllers.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", controllers.GetUsersController)
	user.GET("/:id", controllers.GetUserController)
	user.POST("", controllers.CreateUserController)
	user.PUT("/:id", controllers.UpdateUserController)
	user.DELETE("/:id", controllers.DeleteUserController)

	// book collection
	book := e.Group("/books")
	book.GET("", controllers.GetBookController)
	book.GET("/:id", controllers.GetBookController)
	book.POST("", controllers.CreateBookController)
	book.PUT("/:id", controllers.UpdateBookController)
	book.DELETE("/:id", controllers.DeleteBookController)

	// book collection
	blog := e.Group("/blogs")
	blog.GET("", controllers.GetBlogController)
	blog.GET("/:id", controllers.GetBlogController)
	blog.POST("", controllers.CreateBlogController)
	blog.PUT("/:id", controllers.UpdateBlogController)
	blog.DELETE("/:id", controllers.DeleteBlogController)

}
