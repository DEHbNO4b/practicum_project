package domain

import (
	"fmt"
)

type User struct {
	id       int
	login    string
	password string
	balance  int
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
func (u *User) Id() int {
	return u.id
}

func (u *User) Login() string {
	return u.login
}
func (u *User) Password() string {
	return u.password
}
func (u *User) Balance() int {
	return u.balance
}

func (u *User) SetId(i int) {
	u.id = i
}
func (u *User) SetLogin(l string) {
	u.login = l
}
func (u *User) SetPassword(p string) {
	u.password = p
}
func (u *User) SetBalance(b int) {
	u.balance = b
}
