package debit

import "time"

type Debit struct {
	Order   int       `json:"order"`
	Sum     int       `json:"sum"`
	Time    time.Time `json:"processed_at"`
	User_id int       `json:"-"`
}
