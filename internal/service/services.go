package service

import (
	"context"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
)

type UserService interface {
	AddUser(ctx context.Context, user *domain.User) error
	// GetUserPassword(ctx context.Context, login string) (string, error)
	CheckPassword(ctx context.Context, user *domain.User) (bool, error)
}
type OrderService interface {
	AddOrder(ctx context.Context, order *domain.Order) error
	GetOrdersById(ctx context.Context, id int) ([]*domain.Order, error)
}
