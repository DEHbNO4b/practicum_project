package order

import (
	"context"
	"net/http"

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
	number, err := readOrderNumber(r.Body)
	oc.services.Order.CheckOrderNumber(number)

}
func (oc *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in getOrders handler")
}
