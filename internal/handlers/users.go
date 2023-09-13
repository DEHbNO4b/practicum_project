package handlers

import (
	"context"
	"net/http"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
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

}
