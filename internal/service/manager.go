package service

import (
	"context"
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
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

type Manager struct {
	User    UserService
	Order   OrderService
	Accrual AccrualPointsService
	Balance BalanceService
	Debit   DebitService
}

func NewManager(ctx context.Context, store *storage.Storage) (*Manager, error) {
	if store == nil {
		return nil, errors.New("no store provided")
	}
	return &Manager{
		User:    NewUserWebService(ctx, store),
		Order:   NewOrderWebService(ctx, store),
		Accrual: NewAccrualService(ctx),
		Balance: NewBalanceWebService(ctx, store),
		Debit:   NewDebitWebService(ctx, store),
	}, nil
}
