package order

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

func readNumber(r io.Reader) (int, error) {

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		logger.Log.Error("unable to read from r.Body", zap.Error(err))
		return 0, fmt.Errorf("%s %w", "unable to read order number from r.Body", err)
	}
	number, err := strconv.Atoi(buf.String())
	if err != nil {
		logger.Log.Error("unable to parse r.Body to int", zap.Error(err))
		return 0, fmt.Errorf("%s %w", "unable to read order number from r.Body", err)

	}
	return number, nil
}
func orderHandlerToDomain(o Order) (*domain.Order, error) {
	return domain.NewOrder(o.Number, o.Status, o.Accrual, o.Uploaded_at, o.User_id)
}
