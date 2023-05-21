package token

import (
	authjwt "github.com/asaringo99/authapi/auth/jwt"
	"github.com/asaringo99/authapi/auth/repository"
	domain "github.com/asaringo99/authapi/domain"
	"github.com/labstack/echo/v4"
)

type TokenUsecase struct {
	repository repository.AuthRepositoryInterface
}

func (usecase *TokenUsecase) RetrieveToken(c echo.Context) (string, error) {
	u := c.FormValue("username")
	username := domain.NewUsername(u)
	userid, err := usecase.repository.TransformUserIdFromUserName(&username)
	if err != nil {
		return "", err
	}
	token, err := authjwt.CreateJwtToken(userid.ToValue(), false)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewTokenUsecase(repository repository.AuthRepositoryInterface) *TokenUsecase {
	return &TokenUsecase{
		repository: repository,
	}
}
