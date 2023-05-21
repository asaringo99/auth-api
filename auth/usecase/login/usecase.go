package login

import (
	"github.com/asaringo99/authapi/auth/entity"
	"github.com/asaringo99/authapi/auth/repository"
	"github.com/asaringo99/authapi/domain"
	"github.com/labstack/echo/v4"
)

type LoginUsecase struct {
	repository repository.AuthRepositoryInterface
}

func (usecase *LoginUsecase) Login(c echo.Context) error {
	if err := usecase.canAuthenticate(c); err != nil {
		return err
	}
	return nil
}

func (usecase *LoginUsecase) canAuthenticate(c echo.Context) error {
	u := c.FormValue("username")
	p := c.FormValue("password")
	username := domain.NewUsername(u)
	password := domain.NewPassword(p)
	user := entity.NewUserinfo(username, password)
	if err := usecase.repository.CanAuthenticateUser(&user); err != nil {
		return err
	}
	return nil
}

func NewLoginUsecase(repository repository.AuthRepositoryInterface) *LoginUsecase {
	return &LoginUsecase{
		repository: repository,
	}
}
