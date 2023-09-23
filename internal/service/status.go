package service

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
