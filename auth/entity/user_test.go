package entity

import (
	"reflect"
	"testing"

	"github.com/asaringo99/authapi/domain"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		username domain.Username
		userid   domain.Userid
		want     User
	}{
		{
			name:     "no diff",
			username: domain.NewUsername("user"),
			userid:   domain.NewUserid(1),
			want: User{
				username: domain.NewUsername("user"),
				userid:   domain.NewUserid(1),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := NewUser(tc.username, tc.userid)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestGetUserId(t *testing.T) {
	tests := []struct {
		name     string
		username domain.Username
		userid   domain.Userid
		want     domain.Userid
	}{
		{
			name:     "no diff",
			username: domain.NewUsername("user"),
			userid:   domain.NewUserid(1),
			want:     domain.NewUserid(1),
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		user := NewUser(tc.username, tc.userid)
		got := user.GetUserid()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestGetUserName(t *testing.T) {
	tests := []struct {
		name     string
		username domain.Username
		userid   domain.Userid
		want     domain.Username
	}{
		{
			name:     "no diff",
			username: domain.NewUsername("user"),
			userid:   domain.NewUserid(1),
			want:     domain.NewUsername("user"),
		},
		// TODO: Add Tests
	}
	for _, tc := range tests {
		user := NewUser(tc.username, tc.userid)
		got := user.GetUsername()
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %v, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
