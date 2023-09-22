package balance

import (
	"context"
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/authorization"
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
	claims, err := authorization.GetClaims(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "unable to read token", http.StatusUnauthorized)
		return
	}
	b, err := bc.services.Balance.GetBalance(r.Context(), claims.UserID)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	hBalance := domainBalanceToHandler(b)
	render.JSON(w, r, hBalance)
}
