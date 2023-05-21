package model

import (
	"github.com/asaringo99/authapi/auth/entity"
	"github.com/asaringo99/authapi/domain"
)

type AuthModelInterface interface {
	ContainUser(user *entity.Userinfo) error
	CanAuthenticateUser(user *entity.Userinfo) error
	RegisterUser(*entity.Userinfo) error
	TransformUserIdFromUserName(*domain.Username) (domain.Userid, error)
}
