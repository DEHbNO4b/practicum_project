package agent

import (
	"time"
)

type Order struct {
	Number     string    `json:"order"`
	Status     string    `json:"status"`
	Accrual    int       `json:"accrual"`
	UploadedAt time.Time `json:"-"`
	UserID     int       `json:"-"`
}
