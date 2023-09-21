package agent

import "github.com/DEHbNO4b/practicum_project/internal/domain"

func orderAgentToDomain(o Order) (*domain.Order, error) {
	return domain.NewOrder(o.Number, o.Status, o.Accrual, o.UploadedAt, o.UserID)
}
func domainToAgentOrder(o *domain.Order) *Order {
	return &Order{
		Number:     o.Number(),
		Status:     o.Status(),
		Accrual:    o.Accrual(),
		UploadedAt: o.UpploadedAt(),
		UserID:     o.UserID(),
	}
}
