package agent

import (
	"time"
)

type status int64

const (
	REGISTERED status = iota
	INVALID
	PROCESSING
	PROCESSED
)

func (s status) String() string {
	switch s {
	case REGISTERED:
		return "REGISTERED"
	case INVALID:
		return "INVALID"
	case PROCESSING:
		return "PROCESSING"
	case PROCESSED:
		return "PROCESSED"
	}
	return "unknown"
}

type Order struct {
	Number     int       `json:"number"`
	Status     string    `json:"status"`
	Accrual    int       `json:"accrual"`
	UploadedAt time.Time `json:"-"`
	UserID     int       `json:"-"`
}
