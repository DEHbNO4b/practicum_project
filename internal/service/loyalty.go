package service

import (
	"context"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/config"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
)

type LoyaltyService struct {
	ctx context.Context
	cfg *config.Config
}

func NewLoyaltyService(ctx context.Context) *LoyaltyService {
	config := config.Get()
	return &LoyaltyService{
		ctx: ctx,
		cfg: config,
	}
}
func (svc *LoyaltyService) GetPointsForOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	o, _ := domain.NewOrder(order.Number(), "UPDATED", 777, time.Now(), 0)
	return o, nil
}
