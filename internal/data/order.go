package data

import "time"

type Order struct {
	number       int
	status       string
	accrual      int
	upploaded_at time.Time
}
