package data

import "github.com/gofrs/uuid"

type User struct {
	id       uuid.UUID
	login    string
	password string
}
