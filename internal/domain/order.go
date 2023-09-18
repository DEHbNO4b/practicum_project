package domain

import (
	"fmt"
	"time"
)

type Order struct {
	number      int
	status      string
	accrual     int
	uploaded_at time.Time
	user_id     int
}

func NewOrder(number int, status string, accrual int, uploaded_at time.Time, user_id int) (*Order, error) {
	if number == 0 {
		return nil, fmt.Errorf("%w:number is required", ErrRequired)
	}
	return &Order{number: number, status: status, accrual: accrual, uploaded_at: uploaded_at, user_id: user_id}, nil
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
	return o.uploaded_at
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
	o.uploaded_at = time
}
func (o *Order) SetUserId(id int) {
	o.user_id = id
}
