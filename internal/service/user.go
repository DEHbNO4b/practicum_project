package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/storage"
)

type UserWebService struct {
	ctx     context.Context
	storage *storage.Storage
}

func NewUserWebService(ctx context.Context, storage *storage.Storage) *UserWebService {
	return &UserWebService{
		ctx:     ctx,
		storage: storage,
	}
}
func (svc *UserWebService) AddUser(ctx context.Context, user *domain.User) error {
	hashedPassword := sha256.Sum256([]byte(user.Password()))
	user.SetLogin(user.Login())
	user.SetPassword(base64.StdEncoding.EncodeToString(hashedPassword[:]))
	return svc.storage.User.AddUser(ctx, user)
}
func (svc *UserWebService) CheckPassword(ctx context.Context, user *domain.User) (bool, error) {
	hashedPassword := sha256.Sum256([]byte(user.Password()))
	pass, err := svc.storage.User.GetUserPassword(ctx, user.Login())
	if err != nil {
		return false, err
	}
	if base64.StdEncoding.EncodeToString(hashedPassword[:]) == pass {
		return true, nil
	} else {
		return false, domain.ErrWrongLoginOrPassword
	}
}
