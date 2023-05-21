package entity

import "github.com/asaringo99/authapi/domain"

type Userinfo struct {
	username domain.Username
	password domain.Password
}

func NewUserinfo(username domain.Username, password domain.Password) Userinfo {
	return Userinfo{
		username: username,
		password: password,
	}
}

func (user *Userinfo) GetUsername() domain.Username {
	return user.username
}

func (user *Userinfo) GetPassword() domain.Password {
	return user.password
}
