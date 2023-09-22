package agent

import (
	"context"
	"net/http"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/config"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/go-chi/render"
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
func (a *AccrualAgent) GetAccrual(number string) chan AccrualResponse {
	respCh := make(chan AccrualResponse)
	defer close(respCh)
	go func(number string) {
		req, err := http.NewRequest(http.MethodGet, a.url+`/api/orders/`+number, nil)
		if err != nil {
			logger.Log.Error("acrual server request err", zap.Error(err))
			respCh <- AccrualResponse{order: domain.Order{}, err: err}
		}
		resp, err := a.client.Do(req)
		if err != nil {
			respCh <- AccrualResponse{order: domain.Order{}, err: err}
		}
		defer resp.Body.Close()
		order := Order{}
		switch resp.StatusCode {
		case 200:
			render.DecodeJSON(resp.Body, &order)
			dOrder, err := orderAgentToDomain(order)
			respCh <- AccrualResponse{order: *dOrder, err: err}
		case 204:
			respCh <- AccrualResponse{order: domain.Order{}, err: domain.ErrNotRegistered}
		case 429:
			respCh <- AccrualResponse{order: domain.Order{}, err: domain.ErrTooManyRequests}
		default:
			respCh <- AccrualResponse{order: domain.Order{}, err: domain.ErrUnexpectedRespStatus}
		}
	}(number)
	return respCh
}
