package balance

import "github.com/DEHbNO4b/practicum_project/internal/domain"

func domainBalanceToHandler(b *domain.Balance) *Balance {
	return &Balance{
		Current:   b.Current(),
		Withdrawn: b.Withdrown(),
		UserID:    b.UserID(),
	}
}
func handlerBalanceToDomain(b *Balance) *domain.Balance {
	balance, _ := domain.NewBalance(b.Current, b.Withdrawn, b.UserID)
	return balance
}
