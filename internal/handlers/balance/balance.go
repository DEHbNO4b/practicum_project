package balance

import (
	"context"
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/service"
	"github.com/go-chi/render"
)

type BalanceController struct {
	ctx      context.Context
	services *service.Manager
}

func NewBalance(ctx context.Context, services *service.Manager) *BalanceController {
	return &BalanceController{ctx: ctx, services: services}
}
func (bc *BalanceController) GetBalance(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in get balance handler")
	reqBalance := &Balance{}
	err := render.DecodeJSON(r.Body, reqBalance)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	// err:=bc.services.
	dBalance := handlerBalanceToDomain(reqBalance)
	b, err := bc.services.Balance.GetBalance(r.Context(), dBalance)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, b)
}
