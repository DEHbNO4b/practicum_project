package service

import (
	"context"

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

	svc.storage.Order.AddOrder(ctx, o)
	return nil
}
func (svc *OrderWebService) GetOrdersById(ctx context.Context, id int) ([]*domain.Order, error) {
	return svc.storage.Order.GetOrdersById(ctx, id)
}
func (svc *OrderWebService) GetOrderByNumber(ctx context.Context, number int) (*domain.Order, error) {
	return svc.storage.Order.GetOrderByNumber(ctx, number)
}
func (svc *OrderWebService) UpdateOrder(ctx context.Context, order *domain.Order) error {
	return svc.storage.Order.UpdateOrder(ctx, order)
}
