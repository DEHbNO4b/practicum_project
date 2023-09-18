package order

import "time"

type Order struct {
	Number      int       `json:"number"`
	Status      string    `json:"status"`
	Accrual     int       `json:"accrual"`
	Uploaded_at time.Time `json:"uploaded_at"`
	User_id     int       `json:"-"`
}
