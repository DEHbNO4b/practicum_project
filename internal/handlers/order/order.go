package order

import (
	"context"
	"errors"
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/authorization"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/service"
	"github.com/go-chi/render"
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
	order, _ := orderHandlerToDomain(Order{Number: number, UserID: claims.UserID})
	// if err != nil {
	// 	http.Error(w, "", http.StatusBadRequest)
	// 	return
	// }
	err = oc.services.AddOrder(r.Context(), order, claims)
	switch {
	case errors.Is(err, domain.ErrAccepted):
		http.Error(w, "", http.StatusAccepted)
		return
	case errors.Is(err, domain.ErrHasBeenUpploaded):
		http.Error(w, "", http.StatusConflict)
		return
	case err != nil:
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
func (oc *OrderController) GetOrders(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in getOrders handler")
	claims := authorization.GetClaims(r.Header.Get("Authorization"))
	o, err := oc.services.Order.GetOrdersByID(r.Context(), claims.UserID)
	orders := make([]*Order, 0, 20)
	for _, el := range o {
		hOrder := domainToHandlerOrder(el)
		orders = append(orders, hOrder)
	}

	switch {
	case err == nil: //code200
		render.JSON(w, r, orders)

	case errors.Is(err, domain.ErrNilData):
		http.Error(w, "", http.StatusNoContent) //code 204
	default:
		http.Error(w, "", http.StatusInternalServerError) //code 500
		return

	}

}
