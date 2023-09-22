package agent

import "github.com/DEHbNO4b/practicum_project/internal/domain"

type AccrualResponse struct {
	order domain.Order
	err   error
}
