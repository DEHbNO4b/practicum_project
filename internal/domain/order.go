package domain

import (
	"time"
	"unicode"
)

type Order struct {
	number     string
	status     string
	accrual    float64
	uploadedAt time.Time
	userID     int
}

func NewOrder(number string, status string, accrual float64, uploadedAt time.Time, userID int) (*Order, error) {
	if number == "" {
		return nil, ErrIncorrectOrderNumber
	}
	for _, el := range number {
		if !unicode.IsDigit(el) {
			return nil, ErrIncorrectOrderNumber
		}
	}
	return &Order{number: number, status: status, accrual: accrual, uploadedAt: uploadedAt, userID: userID}, nil
}
func (o *Order) Number() string {
	return o.number
}
func (o *Order) Status() string {
	return o.status
}
func (o *Order) Accrual() float64 {
	return o.accrual
}
func (o *Order) UpploadedAt() time.Time {
	return o.uploadedAt
}
func (o *Order) UserID() int {
	return o.userID
}
func (o *Order) SetNumber(number string) {
	o.number = number
}
func (o *Order) SetStatus(status string) {
	o.status = status
}
func (o *Order) SetAccrual(accrual float64) {
	o.accrual = accrual
}
func (o *Order) SetTime(time time.Time) {
	o.uploadedAt = time
}
func (o *Order) SetUserID(id int) {
	o.userID = id
}
