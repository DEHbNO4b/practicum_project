package service

import (
	"context"
	"errors"

	"github.com/DEHbNO4b/practicum_project/internal/authorization"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
)

type OrderWebService struct {
	ctx     context.Context
	storage *storage.Storage
}

func NewOrderWebService(ctx context.Context, storage *storage.Storage) *OrderWebService {
	return &OrderWebService{
		ctx:     ctx,
		storage: storage,
	}
}
func (svc *OrderWebService) AddOrder(ctx context.Context, o *domain.Order, claims authorization.Claims) error {
	order, err := svc.storage.Order.GetOrderByNumber(ctx, o.Number())
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return err
	}
	if !errors.Is(err, domain.ErrNotFound) {
		if order.UserId() == claims.UserID {
			return domain.ErrAccepted
		}
		if order.UserId() != claims.UserID {
			return domain.ErrHasBeenUpploaded
		}
	}

	svc.storage.Order.AddOrder(ctx, order)
	return nil
}
func (svc *OrderWebService) GetOrdersById(ctx context.Context, id int) ([]*domain.Order, error) {
	return svc.storage.Order.GetOrdersById(ctx, id)
}
