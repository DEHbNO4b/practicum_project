package agent

import (
	"time"
)

type Order struct {
	Number     string    `json:"order"`
	Status     string    `json:"status"`
	Accrual    float64   `json:"accrual"`
	UploadedAt time.Time `json:"-"`
	UserID     int       `json:"-"`
}
