package service

import (
	"context"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
)

type BalanceWebService struct {
	ctx     context.Context
	storage *storage.Storage
}

func NewBalanceWebService(ctx context.Context, storage *storage.Storage) *BalanceWebService {
	return &BalanceWebService{
		ctx:     ctx,
		storage: storage,
	}
}
func (svc *BalanceWebService) GetBalance(ctx context.Context, id int) (*domain.Balance, error) {
	b, err := svc.storage.Balance.GetById(ctx, id)
	return b, err
}
