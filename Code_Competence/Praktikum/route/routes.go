package route

import (
	"code_competence/constant"
	"code_competence/controller"
	"code_competence/repository/database"
	"code_competence/usecase"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := database.NewUserRepository(db)
	itemRepository := database.NewItemRepository(db)
	categoryRepository := database.NewCategoryRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	itemUsecase := usecase.NewItemUseCase(itemRepository)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)

	userController := controller.NewUserController(userUsecase)
	itemController := controller.NewItemController(itemUsecase)
	categoryController := controller.NewCategoryController(categoryUsecase)

	e.Validator = &customValidator{validator: validator.New()}

	e.POST("/register", userController.RegisterUserController)
	e.POST("/login", userController.LoginUserController)

	items := e.Group("/items", middleware.JWT([]byte(constant.SECRET_JWT)))
	items.POST("", itemController.CreateItemController)
	items.GET("", itemController.GetAllItemsController)
	items.GET("/id/:id", itemController.GetItemByIDController)
	items.PUT("/id/:id", itemController.UpdateItemByIDController)
	items.DELETE("/id/:id", itemController.DeleteItemByIDController)
	items.POST("/category", categoryController.CreateCategoryController)
	items.GET("/category/:category_id", categoryController.GetAllByCategoryIDController)
	items.GET("/name", itemController.GetItemByItemNameController)
}
