package usecase

import (
	"code_competence/middleware"
	"code_competence/model"
	"code_competence/model/payload"
	"code_competence/repository/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	CreateUser(req *payload.CreateUserRequest) (*model.User, error)
	LoginUser(email, password string) (*model.User, error)
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepo database.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) CreateUser(req *payload.CreateUserRequest) (*model.User, error) {
	_, err := u.userRepository.GetUserByEmail(req.Email)
	if err == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email Already Registered")
	}
	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err = u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) LoginUser(email, password string) (*model.User, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email Not Registered")
	}
	if user.Password != password {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid Password")
	}

	token, err := middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Error Create Token")
	}
	user.Token = token
	return user, nil
}
