package service

import (
	"context"
	"errors"

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
func (svc *OrderWebService) AddOrder(ctx context.Context, o *domain.Order, id int) error {
	order, err := svc.GetOrderByNumber(ctx, o.Number())
	if err != nil {
		if !errors.Is(err, domain.ErrNotFound) {
			return err
		}
		return svc.storage.Order.AddOrder(ctx, o)
	}
	if order.UserID() == o.UserID() {
		return domain.ErrAccepted
	} else {
		return domain.ErrConflict
	}
	return nil

}
func (svc *OrderWebService) GetOrdersByID(ctx context.Context, id int) ([]*domain.Order, error) {
	return svc.storage.Order.GetOrdersByID(ctx, id)
}
func (svc *OrderWebService) GetOrderByNumber(ctx context.Context, number int) (*domain.Order, error) {
	return svc.storage.Order.GetOrderByNumber(ctx, number)
}
func (svc *OrderWebService) UpdateOrder(ctx context.Context, order *domain.Order) error {
	return svc.storage.Order.UpdateOrder(ctx, order)
}
