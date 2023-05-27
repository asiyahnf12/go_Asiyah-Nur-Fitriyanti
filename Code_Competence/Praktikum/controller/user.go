package controller

import (
	"code_competence/model/payload"
	"code_competence/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	RegisterUserController(c echo.Context) error
	LoginUserController(c echo.Context) error
}

type userController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{userUsecase: userUsecase}
}

func (u *userController) RegisterUserController(c echo.Context) error {
	req := payload.CreateUserRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	user, err := u.userUsecase.CreateUser(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to register user",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
		"user":    user,
	})
}

func (a *userController) LoginUserController(c echo.Context) error {
	payload := payload.LoginRequest{}
	c.Bind(&payload)

	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	user, err := a.userUsecase.LoginUser(payload.Email, payload.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}
