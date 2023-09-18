package service

import (
	"context"

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
func (svc *OrderWebService) AddOrder(ctx context.Context, order *domain.Order) error {
	return svc.storage.Order.AddOrder(ctx, order)
}
func (svc *OrderWebService) GetOrdersById(ctx context.Context, id int) ([]*domain.Order, error) {
	return svc.storage.Order.GetOrdersById(ctx, id)
}
