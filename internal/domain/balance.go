package domain

import "fmt"

type Balance struct {
	current   int
	withdrown int
	user_id   int
}

func NewBalance(c, w, id int) (*Balance, error) {
	if id == 0 {
		return nil, fmt.Errorf("%w :id is required", ErrRequired)
	}
	return &Balance{current: c, withdrown: w, user_id: id}, nil
}

func (b *Balance) Current() int {
	return b.current
}
func (b *Balance) Withdrown() int {
	return b.withdrown
}
func (b *Balance) User_id() int {
	return b.user_id
}
