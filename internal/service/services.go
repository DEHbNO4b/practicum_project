package service

import (
	"context"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
)

type UserService interface {
	AddUser(ctx context.Context, user *domain.User) (int64, error)
	GetUser(ctx context.Context, user *domain.User) (*domain.User, error)
	CheckPassword(ctx context.Context, user *domain.User) (bool, error)
}
type OrderService interface {
	AddOrder(ctx context.Context, order *domain.Order, id int) error
	UpdateOrder(ctx context.Context, order *domain.Order) error
	GetOrdersByID(ctx context.Context, id int) ([]*domain.Order, error)
	GetOrderByNumber(ctx context.Context, number string) (*domain.Order, error)
}
type AccrualPointsService interface {
	GetPointsForOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
}
type BalanceService interface {
	GetBalance(ctx context.Context, id int) (*domain.Balance, error)
}
type DebitService interface {
	AddDebit(ctx context.Context, debit *domain.Debit) error
	GetDebitsByID(ctx context.Context, id int) ([]*domain.Debit, error)
}
