package signup

import (
	"github.com/asaringo99/authapi/auth/entity"
	"github.com/asaringo99/authapi/auth/repository"
	"github.com/asaringo99/authapi/domain"
	"github.com/labstack/echo/v4"
)

type SignUpUsecase struct {
	repository repository.AuthRepositoryInterface
}

func (usecase *SignUpUsecase) RegisterUser(c echo.Context) error {
	u := c.FormValue("username")
	p := c.FormValue("password")
	username := domain.NewUsername(u)
	password := domain.NewHashedPassword(p)
	user := entity.NewUserinfo(username, password)
	if err := usecase.repository.ContainUser(&user); err != nil {
		return err
	}
	usecase.repository.Register(&user)
	return nil
}

func NewSignUpUsecase(repository repository.AuthRepositoryInterface) *SignUpUsecase {
	return &SignUpUsecase{
		repository: repository,
	}
}
