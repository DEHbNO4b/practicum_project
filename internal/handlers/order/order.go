package order

import (
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/logger"
)

type Gopher struct {
}

func (g *Gopher) Calculate(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in Calculate handler")
}
func (g *Gopher) GetOrder(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in getOrders handler")
}
