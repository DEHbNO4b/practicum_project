package debit

import "time"

type Debit struct {
	Order  string    `json:"order"`
	Sum    float64   `json:"sum"`
	Time   time.Time `json:"processed_at"`
	UserID int       `json:"-"`
}
