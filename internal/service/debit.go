package service

import (
	"context"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
)

type DebitWebService struct {
	ctx     context.Context
	storage *storage.Storage
}

func NewDebitWebService(ctx context.Context, storage *storage.Storage) *DebitWebService {
	return &DebitWebService{
		ctx:     ctx,
		storage: storage,
	}
}
func (svc *DebitWebService) AddDebit(ctx context.Context, debit *domain.Debit) error {
	logger.Log.Info("in AddDebit() at DebitWebService{}")
	err := svc.storage.Debit.AddDebit(ctx, debit)
	return err
}
func (svc *DebitWebService) GetDebitsById(ctx context.Context, id int) ([]*domain.Debit, error) {
	logger.Log.Info("in GetDebitsById() at DebitWebService{}")
	d, err := svc.storage.Debit.GetDebitsById(ctx, id)
	return d, err
}
