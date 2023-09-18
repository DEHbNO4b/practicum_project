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
}
type OrderRepository interface {
	AddOrder(ctx context.Context, order *domain.Order) error
	GetOrdersById(ctx context.Context, id int) ([]*domain.Order, error)
	GetOrderByNumber(ctx context.Context, number int) (*domain.Order, error)
}

type Storage struct {
	User  UserRepository
	Order OrderRepository
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
	if udb != nil {
		store.User = udb
	}
	if odb != nil {
		store.Order = odb
	}

	return store, nil
}
