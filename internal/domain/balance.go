package domain

import "fmt"

type Balance struct {
	current   int
	withdrown int
	userID    int
}

func NewBalance(c, w, id int) (*Balance, error) {
	if id == 0 {
		return nil, fmt.Errorf("%w :id is required", ErrRequired)
	}
	return &Balance{current: c, withdrown: w, userID: id}, nil
}

func (b *Balance) Current() int {
	return b.current
}
func (b *Balance) Withdrown() int {
	return b.withdrown
}
func (b *Balance) UserID() int {
	return b.userID
}
