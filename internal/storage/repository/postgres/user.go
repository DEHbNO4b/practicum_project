package postgres

import (
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
)

type User struct {
	Id       int
	Login    string
	Password string
	Balance  int
}

func userStoreToDomain(u *User) (*domain.User, error) {
	if u == nil {
		return nil, errors.New("store user is nil")
	}
	return domain.NewUser(u.Id, u.Login, u.Password, u.Balance)
}
func userDomainToStore(u *domain.User) *User {
	return &User{Login: u.Login(), Password: u.Password()}
}
