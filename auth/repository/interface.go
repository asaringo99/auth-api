package repository

import (
	"github.com/asaringo99/authapi/auth/entity"
	"github.com/asaringo99/authapi/domain"
)

type AuthRepositoryInterface interface {
	ContainUser(*entity.Userinfo) error
	CanAuthenticateUser(*entity.Userinfo) error
	Register(*entity.Userinfo) error
	TransformUserIdFromUserName(*domain.Username) (*domain.Userid, error)
}
