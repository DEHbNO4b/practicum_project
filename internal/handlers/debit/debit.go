package debit

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

type DebitController struct {
	ctx      context.Context
	services *service.Manager
}

func NewDebit(ctx context.Context, services *service.Manager) *DebitController {
	return &DebitController{ctx: ctx, services: services}
}
func (dc *DebitController) AddDebit(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in add debit handler")
	debit := Debit{}
	err := render.DecodeJSON(r.Body, &debit)
	if err != nil {
		logger.Log.Error("unable to decode json from r.Body()")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	claims, err := authorization.GetClaims(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "unable to read token", http.StatusUnauthorized)
		return
	}
	domainDebit := handlerDebitToDomain(&debit)
	domainDebit.SetID(claims.UserID)
	err = dc.services.AddDebit(r.Context(), domainDebit)
	switch {
	case errors.Is(err, domain.ErrIncorrectOrderNumber):
		http.Error(w, "", http.StatusUnprocessableEntity) //status 422
		return
	case errors.Is(err, domain.ErrNotEnaugh):
		http.Error(w, "", http.StatusPaymentRequired) //status 402
		return
	case err != nil:
		http.Error(w, "", http.StatusInternalServerError) //status 500
		return
	}
	w.WriteHeader(http.StatusOK) //status 200

}
func (dc *DebitController) GetAllDebits(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in add debit handler")
	claims, err := authorization.GetClaims(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "unable to read token", http.StatusUnauthorized)
		return
	}
	debits, err := dc.services.Debit.GetDebitsByID(r.Context(), claims.UserID)
	switch {
	case errors.Is(err, domain.ErrNotFound):
		http.Error(w, "", http.StatusNoContent) //204
		return
	case err != nil:
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	hDebits := make([]*Debit, 0, 10)
	for _, el := range debits {
		hDebits = append(hDebits, domainDebitToHandler(el))
	}
	render.JSON(w, r, hDebits)
}
