package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

type UserService interface {
	AddUser(ctx context.Context, user *domain.User) error
}
type UserRegister struct {
	userRepo UserService
}

func NewRegister(userRepo UserService) *UserRegister {
	r := UserRegister{userRepo: userRepo}
	return &r
}
func (u *UserRegister) Register(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in Register handler")
	user, err := readUser(r.Context(), r.Body)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.Log.Sugar().Infof("%+v", user)
	dUser, err := userHandlerToDomain(user)
	if err != nil {
		logger.Log.Error("unable to create user", zap.Error(err))
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = u.userRepo.AddUser(r.Context(), dUser)
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
func (u *UserRegister) Login(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("in Login handler")

}
