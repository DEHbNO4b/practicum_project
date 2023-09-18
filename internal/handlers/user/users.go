package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/authorization"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/service"
	"go.uber.org/zap"
)

type UserService interface {
	AddUser(ctx context.Context, user *domain.User) error
	CheckPassword(ctx context.Context, user *domain.User) (bool, error)
}
type UserController struct {
	ctx      context.Context
	services *service.Manager
}

func NewUsers(ctx context.Context, services *service.Manager) *UserController {
	r := UserController{ctx: ctx, services: services}
	return &r
}
func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in Register handler")
	user, err := readUser(r.Context(), r.Body)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dUser, err := userHandlerToDomain(user)
	if err != nil {
		logger.Log.Error("unable to create user", zap.Error(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = uc.services.User.AddUser(r.Context(), dUser)
	if err != nil {
		if errors.Is(err, domain.ErrUniqueViolation) {
			http.Error(w, "", http.StatusConflict)
			return
		} else {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in Login handler")
	user, err := readUser(r.Context(), r.Body)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dUser, err := userHandlerToDomain(user)
	if err != nil {
		logger.Log.Error("unable to create user", zap.Error(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	checked, err := uc.services.User.CheckPassword(r.Context(), dUser)
	if err != nil {
		if errors.Is(err, domain.ErrWrongLoginOrPassword) {
			http.Error(w, "", http.StatusUnauthorized)
		}
	}
	if checked {
		jwt, err := authorization.BuildJWTString()
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		w.Header().Set("Authorization", jwt)
		w.Write([]byte("password is correct"))
	} else {
		http.Error(w, "password not correct", http.StatusUnauthorized)
	}

}
