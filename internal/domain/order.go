package domain

import (
	"fmt"
	"time"
)

type Order struct {
	number       int
	status       string
	accrual      int
	upploaded_at time.Time
	user_id      int
}

func NewOrder(number int) (*Order, error) {
	if number == 0 {
		return nil, fmt.Errorf("%w:number is required", ErrRequired)
	}
	return &Order{number: number}, nil
}
func (o *Order) Number() int {
	return o.number
}
func (o *Order) Status() string {
	return o.status
}
func (o *Order) Accrual() int {
	return o.accrual
}
func (o *Order) UpploadedAt() time.Time {
	return o.upploaded_at
}
func (o *Order) UserId() int {
	return o.user_id
}
func (o *Order) SetStatus(status string) {
	o.status = status
}
func (o *Order) SetAccrual(accrual int) {
	o.accrual = accrual
}
func (o *Order) SetTime(time time.Time) {
	o.upploaded_at = time
}
func (o *Order) SetUserId(id int) {
	o.user_id = id
}
