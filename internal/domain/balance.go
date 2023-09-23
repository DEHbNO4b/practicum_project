package domain

import "fmt"

type Balance struct {
	current   float64
	withdrown float64
	userID    int
}

func NewBalance(c, w float64, id int) (*Balance, error) {
	if id == 0 {
		return nil, fmt.Errorf("%w :id is required", ErrRequired)
	}
	return &Balance{current: c, withdrown: w, userID: id}, nil
}

func (b *Balance) Current() float64 {
	return b.current
}
func (b *Balance) Withdrown() float64 {
	return b.withdrown
}
func (b *Balance) UserID() int {
	return b.userID
}
func (b *Balance) SetCurrent(c float64) {
	b.current = c
}
func (b *Balance) SetWithdrown(w float64) {
	b.withdrown = w
}

// func (b *Balance) AddToCurrent(sum float64) {
// 	b.current += sum
// }
// func (b *Balance) WriteOff(sum float64) error {
// 	if sum > b.current {
// 		return ErrNotEnaugh
// 	}
// 	b.current -= sum
// 	return nil
// }
