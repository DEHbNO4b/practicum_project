package postgres

import (
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
)

type User struct {
	Login    string
	Password string
}

func userStoreToDomain(u *User) (*domain.User, error) {
	if u == nil {
		return nil, errors.New("store user is nil")
	}
	return domain.NewUser(u.Login, u.Password)
}
func userDomainToStore(u *domain.User) *User {
	return &User{Login: u.Login(), Password: u.Password()}
}
