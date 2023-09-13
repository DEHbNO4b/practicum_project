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
}

func NewOrder(number, accrual int, status string, time time.Time) (*Order, error) {
	if number == 0 {
		return nil, fmt.Errorf("%w:number is required", ErrRequired)
	}
	// if accrual == 0 {
	// 	return nil, fmt.Errorf("%w:accrual is required", ErrRequired)
	// }
	// if status == "" {
	// 	return nil, fmt.Errorf("%w:status is required", ErrRequired)
	// }
	// if time.IsZero() {
	// 	return nil, fmt.Errorf("%w:time is required", ErrRequired)
	// }
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
func (o *Order) SetStatus(status string) {
	o.status = status
}
func (o *Order) SetAccrual(accrual int) {
	o.accrual = accrual
}
func (o *Order) SetTime(time time.Time) {
	o.upploaded_at = time
}
