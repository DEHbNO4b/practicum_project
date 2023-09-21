package domain

import (
	"time"
	"unicode"
)

type Debit struct {
	order  string
	sum    float64
	time   time.Time
	userID int
}

func NewDebit(orderNumber string, sum float64, t time.Time, id int) (*Debit, error) {
	if orderNumber == "" {
		return nil, ErrIncorrectOrderNumber
	}
	for _, el := range orderNumber {
		if !unicode.IsDigit(el) {
			return nil, ErrIncorrectOrderNumber
		}
	}
	return &Debit{
		order:  orderNumber,
		sum:    sum,
		time:   t,
		userID: id,
	}, nil
}
func (d *Debit) Order() string {
	return d.order
}
func (d *Debit) Sum() float64 {
	return d.sum
}
func (d *Debit) Time() time.Time {
	return d.time
}
func (d *Debit) UserID() int {
	return d.userID
}
func (d *Debit) SetOrder(o string) {
	d.order = o
}
func (d *Debit) SetSum(s float64) {
	d.sum = s
}
func (d *Debit) SetTime(t time.Time) {
	d.time = t
}
func (d *Debit) SetID(id int) {
	d.userID = id
}
