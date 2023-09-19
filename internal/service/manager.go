package service

import (
	"context"
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/authorization"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
)

type Manager struct {
	User    UserService
	Order   OrderService
	Loyalty LoyaltyPointsService
}

func NewManager(ctx context.Context, store *storage.Storage) (*Manager, error) {
	if store == nil {
		return nil, errors.New("no store provided")
	}
	return &Manager{
		User:    NewUserWebService(ctx, store),
		Order:   NewOrderWebService(ctx, store),
		Loyalty: NewLoyaltyService(ctx),
	}, nil
}
func (m *Manager) AddOrder(ctx context.Context, o *domain.Order, claims authorization.Claims) error {
	order, err := m.Order.GetOrderByNumber(ctx, o.Number())
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		//internal err
		return err
	}
	if err == nil {
		if order.UserId() == claims.UserID {
			return domain.ErrAccepted
		}
		if order.UserId() != claims.UserID {
			return domain.ErrHasBeenUpploaded
		}
	}
	err = m.Order.AddOrder(ctx, o, claims)
	if err != nil {
		return err
	}
	lOrder, err := m.Loyalty.GetPointsForOrder(ctx, o)
	if err != nil {
		return err
	}
	err = m.Order.UpdateOrder(ctx, lOrder)
	if err != nil {
		//internal err
		return err
	}
	return nil
}
