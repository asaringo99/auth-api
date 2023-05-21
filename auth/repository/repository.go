package repository

import (
	"github.com/asaringo99/authapi/auth/entity"
	"github.com/asaringo99/authapi/auth/model"
	domain "github.com/asaringo99/authapi/domain"
)

type AuthRepositoryImpl struct {
	authModel model.AuthModelInterface
}

func (r *AuthRepositoryImpl) ContainUser(user *entity.Userinfo) error {
	err := r.authModel.ContainUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepositoryImpl) CanAuthenticateUser(user *entity.Userinfo) error {
	if err := r.authModel.CanAuthenticateUser(user); err != nil {
		return err
	}
	return nil
}

func (r *AuthRepositoryImpl) Register(user *entity.Userinfo) error {
	if err := r.authModel.RegisterUser(user); err != nil {
		return err
	}
	return nil
}

func (r *AuthRepositoryImpl) TransformUserIdFromUserName(u *domain.Username) (*domain.Userid, error) {
	userid, err := r.authModel.TransformUserIdFromUserName(u)
	if err != nil {
		return nil, err
	}
	return &userid, nil
}

func NewAuthRepository(m model.AuthModelInterface) AuthRepositoryInterface {
	return &AuthRepositoryImpl{
		authModel: m,
	}
}
