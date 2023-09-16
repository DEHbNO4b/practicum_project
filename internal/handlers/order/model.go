package order

import "time"

type Order struct {
	Number       int
	Status       string
	Accrual      int
	Upploaded_at time.Time
}
