package agent

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/config"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

type AccrualAgent struct {
	ctx    context.Context
	client http.Client
	url    string
}

func NewAccrualAgent(ctx context.Context) *AccrualAgent {
	cfg := config.Get()
	cl := http.Client{Timeout: 1 * time.Second}
	a := &AccrualAgent{ctx: ctx, client: cl, url: cfg.AccrualSystemAdress}
	return a
}
func (a *AccrualAgent) GetAccrual(number int) (*domain.Order, error) {

	num := strconv.Itoa(number)
	req, err := http.NewRequest(http.MethodGet, a.url+`/api/orders/`+num, nil)
	if err != nil {
		logger.Log.Error("acrual server request err", zap.Error(err))
		return nil, err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// if err != nil {
	// 	logger.Log.Error("acrual server returned err", zap.Error(err))
	// 	return nil, err
	// }
	switch resp.StatusCode {
	case 200:
	case 204:
	case 429:
	}
	return domain.NewOrder(number, "", 0, time.Now(), 0)
}
