package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/authorization"
	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/DEHbNO4b/practicum_project/internal/service"
	"github.com/go-chi/render"
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
	// user, err := readUser(r.Context(), r.Body)
	user := User{}
	err := render.DecodeJSON(r.Body, user)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest) //400
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
	requestUser, err := readUser(r.Context(), r.Body)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	dUser, err := userHandlerToDomain(requestUser)
	if err != nil {
		logger.Log.Error("unable to create user", zap.Error(err))
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	_, err = uc.services.User.CheckPassword(r.Context(), dUser)

	switch {
	case errors.Is(err, domain.ErrNotFound):
		http.Error(w, "wrong login or password", http.StatusUnauthorized)
		return
	case err != nil:
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	// if !isCorrect || errors.Is(err, domain.ErrWrongLoginOrPassword) {
	// 	http.Error(w, "", http.StatusUnauthorized)
	// 	return
	// } else if err != nil {
	// 	http.Error(w, "", http.StatusInternalServerError)
	// 	return
	// }
	userFromDB, err := uc.services.User.GetUser(r.Context(), dUser)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	jwt, err := authorization.BuildJWTString(userFromDB.ID())
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", jwt)
	w.Write([]byte("password is correct"))

}
