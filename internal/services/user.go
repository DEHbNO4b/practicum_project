package services

import (
	"context"
	"crypto/sha256"
	"encoding/base64"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
)

type UserRepository interface {
	AddUser(ctx context.Context, user *domain.User) error
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		userRepo: repo,
	}
}
func (us *UserService) AddUser(ctx context.Context, user *domain.User) error {
	hashedPassword := sha256.Sum256([]byte(user.Password()))
	user.SetLogin(user.Login())
	user.SetPassword(base64.StdEncoding.EncodeToString(hashedPassword[:]))
	return us.userRepo.AddUser(ctx, user)
}
