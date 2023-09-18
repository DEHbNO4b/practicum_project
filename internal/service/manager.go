package service

import (
	"context"
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/storage"
)

type Manager struct {
	User  UserService
	Order OrderService
}

func NewManager(ctx context.Context, store *storage.Storage) (*Manager, error) {
	if store == nil {
		return nil, errors.New("no store provided")
	}
	return &Manager{
		User:  NewUserWebService(ctx, store),
		Order: NewOrderWebService(ctx, store),
	}, nil
}
