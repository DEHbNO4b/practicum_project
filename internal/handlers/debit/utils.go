package debit

import "github.com/DEHbNO4b/practicum_project/internal/domain"

func domainDebitToHandler(d *domain.Debit) *Debit {
	return &Debit{
		Order:  d.Order(),
		Sum:    d.Sum(),
		Time:   d.Time(),
		UserID: d.UserID(),
	}
}
func handlerDebitToDomain(d *Debit) *domain.Debit {
	Debit, _ := domain.NewDebit(d.Order, d.Sum, d.Time, d.UserID)
	return Debit
}
