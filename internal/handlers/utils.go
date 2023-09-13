package handlers

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/DEHbNO4b/practicum_project/internal/domain"
)

var ErrNil = errors.New("nil data")

func readUser(ctx context.Context, r io.Reader) (domain.User, error) {
	buf := new(bytes.Buffer)
	n, err := buf.ReadFrom(r)
	if n == 0 {
		return domain.User{}, ErrNil
	}
	if err != nil {
		return domain.User{}, err
	}
	user := domain.User{}
	// err := json.Unmarshal(, user)
}
