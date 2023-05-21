package entity

import "github.com/asaringo99/authapi/domain"

type User struct {
	username domain.Username
	userid   domain.Userid
}

func NewUser(username domain.Username, userid domain.Userid) User {
	return User{
		username: username,
		userid:   userid,
	}
}

func (user *User) GetUsername() domain.Username {
	return user.username
}

func (user *User) GetUserid() domain.Userid {
	return user.userid
}
