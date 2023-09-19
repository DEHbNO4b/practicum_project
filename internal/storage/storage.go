package storage

import (
	"context"
	"fmt"

	"github.com/DEHbNO4b/practicum_project/internal/config"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/storage/repository/postgres"
)

type UserRepository interface {
	AddUser(ctx context.Context, user *domain.User) error
	GetUser(ctx context.Context, login string) (*domain.User, error)
	Close()
}
type OrderRepository interface {
	AddOrder(ctx context.Context, order *domain.Order) error
	UpdateOrder(ctx context.Context, order *domain.Order) error
	GetOrdersById(ctx context.Context, id int) ([]*domain.Order, error)
	GetOrderByNumber(ctx context.Context, number int) (*domain.Order, error)
	Close()
}
type BalanceRepository interface {
	GetById(ctx context.Context, id int) (*domain.Balance, error)
	UpdateBalance(ctx context.Context, balance *domain.Balance) error
	Close()
}
type DebitRepository interface {
	AddDebit(ctx context.Context, d *domain.Debit) error
	GetDebitsById(ctx context.Context, id int) ([]*domain.Debit, error)
}
type Storage struct {
	User    UserRepository
	Order   OrderRepository
	Balance BalanceRepository
	Debit   DebitRepository
}

func New(ctx context.Context) (*Storage, error) {
	cfg := config.Get()
	store := &Storage{}
	//connect to user_db
	udb, err := postgres.NewUserDB(cfg.Database_url)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "unable to create postgres user_DB", err)
	}
	odb, err := postgres.NewOrderDB(cfg.Database_url)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "unable to create postgres order_DB", err)
	}
	bdb, err := postgres.NewBalanceDB(cfg.Database_url)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "unable to create postgres balance_DB", err)
	}
	ddb, err := postgres.NewDebitDB(cfg.Database_url)
	if err != nil {
		return nil, fmt.Errorf("%s %w", "unable to create postgres debit_DB", err)
	}
	if udb != nil {
		store.User = udb
	}
	if odb != nil {
		store.Order = odb
	}
	if bdb != nil {
		store.Balance = bdb
	}
	if ddb != nil {
		store.Debit = ddb
	}
	return store, nil
}
func (s *Storage) Close() {
	s.User.Close()
	s.Order.Close()
	s.Balance.Close()
}
