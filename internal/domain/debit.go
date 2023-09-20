package domain

import (
	"fmt"
	"time"
)

type Debit struct {
	order  int
	sum    int
	time   time.Time
	userID int
}

func NewDebit(o int, sum int, t time.Time, id int) (*Debit, error) {
	if o == 0 {
		return nil, fmt.Errorf("%w: order number is required", ErrRequired)
	}
	return &Debit{
		order:  o,
		sum:    sum,
		time:   t,
		userID: id,
	}, nil
}
func (d *Debit) Order() int {
	return d.order
}
func (d *Debit) Sum() int {
	return d.sum
}
func (d *Debit) Time() time.Time {
	return d.time
}
func (d *Debit) UserID() int {
	return d.userID
}
func (d *Debit) SetOrder(o int) {
	d.order = o
}
func (d *Debit) SetSum(s int) {
	d.sum = s
}
func (d *Debit) SetTime(t time.Time) {
	d.time = t
}
func (d *Debit) SetID(id int) {
	d.userID = id
}
