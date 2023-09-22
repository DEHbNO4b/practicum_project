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
	b, err := svc.storage.Balance.GetByID(ctx, id)
	return b, err
}
func (svc *BalanceWebService) NewBalance(ctx context.Context, id int) error {
	b, err := domain.NewBalance(0, 0, id)
	if err != nil {
		return err
	}
	err = svc.storage.Balance.NewBalance(ctx, b)
	if err != nil {
		return err
	}
	return nil
}

func (svc *BalanceWebService) Withdrawm(ctx context.Context, sum float64, id int) error {

	b, err := svc.storage.Balance.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if b.Current() < sum {
		return domain.ErrNotEnaugh
	}
	b.SetCurrent(b.Current() - sum)
	err = svc.storage.Balance.UpdateBalance(ctx, b)
	if err != nil {
		return err
	}

	return nil
}
