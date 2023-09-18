package order

import (
	"context"
	"errors"
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/authorization"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/service"
)

type OrderController struct {
	ctx      context.Context
	services *service.Manager
}

func NewOrders(ctx context.Context, services *service.Manager) *OrderController {
	return &OrderController{ctx: ctx, services: services}
}
func (oc *OrderController) Calculate(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in Calculate handler")
	number, err := readNumber(r.Body)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	claims := authorization.GetClaims(r.Header.Get("Authorization"))
	order, _ := orderHandlerToDomain(Order{Number: number, User_id: claims.UserID})
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	err = oc.services.Order.AddOrder(r.Context(), order, claims)
	switch {
	case errors.Is(err, domain.ErrAccepted):
		http.Error(w, "", http.StatusAccepted)
		return
	case errors.Is(err, domain.ErrHasBeenUpploaded):
		http.Error(w, "", http.StatusConflict)
		return
	}

}
func (oc *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in getOrders handler")
}
