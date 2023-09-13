package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"go.uber.org/zap"
)

var ErrNil = errors.New("nil data")

func readUser(ctx context.Context, r io.Reader) (User, error) {
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(r)
	if n == 0 {
		return User{}, ErrNil
	}
	if err != nil {
		return User{}, err
	}
	user := User{}
	err = json.Unmarshal(buf.Bytes(), &user)
	if err != nil {
		logger.Log.Error("unable to unmarshal json", zap.Error(err))
		return User{}, err
	}
	return user, nil
}
func userHandlerToDomain(user User) (*domain.User, error) {
	return domain.NewUser(user.Login, user.Password)
}
