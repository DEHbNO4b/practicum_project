package debit

import "github.com/DEHbNO4b/practicum_project/internal/domain"

func domainDebitToHandler(d *domain.Debit) *Debit {
	return &Debit{
		Order:   d.Order(),
		Sum:     d.Sum(),
		Time:    d.Time(),
		User_id: d.UserId(),
	}
}
func handlerDebitToDomain(d *Debit) *domain.Debit {
	Debit, _ := domain.NewDebit(d.Order, d.Sum, d.Time, d.User_id)
	return Debit
}
