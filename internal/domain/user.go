package domain

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type User struct {
	id       uuid.UUID
	login    string
	password string
}

func NewUser(login, password string) (*User, error) {
	if login == "" {
		return nil, fmt.Errorf("%w:login is required", ErrRequired)
	}
	if password == "" {
		return nil, fmt.Errorf("%w:password is required", ErrRequired)
	}
	return &User{login: login, password: password}, nil
}
func (u *User) Login() string {
	return u.login
}
func (u *User) Password() string {
	return u.password
}
