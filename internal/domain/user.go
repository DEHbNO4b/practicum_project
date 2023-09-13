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

func NewUser(id uuid.UUID, login, password string) (*User, error) {
	if id == uuid.Nil {
		return nil, fmt.Errorf("%w:id is required", ErrRequired)
	}
	if login == "" {
		return nil, fmt.Errorf("%w:login is required", ErrRequired)
	}
	if password == "" {
		return nil, fmt.Errorf("%w:password is required", ErrRequired)
	}
	return &User{id: id, login: login, password: password}, nil
}
