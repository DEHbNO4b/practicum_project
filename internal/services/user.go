package services

import (
	"context"

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
	return us.userRepo.AddUser(ctx, user)
}
