package service

import (
	"context"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/agent"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
)

type AccrualService struct {
	ctx   context.Context
	agent *agent.AccrualAgent
}

func NewAccrualService(ctx context.Context) *AccrualService {
	ag := agent.NewAccrualAgent(ctx)
	return &AccrualService{
		ctx:   ctx,
		agent: ag,
	}
}
func (svc *AccrualService) GetPointsForOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	o, _ := domain.NewOrder(order.Number(), "UPDATED", 777, time.Now(), 0)

	return o, nil
}
