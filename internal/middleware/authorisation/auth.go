package authorisation

import (
	"net/http"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/logger"
	"github.com/golang-jwt/jwt/v4"
)

const TOKEN_EXP = time.Hour * 3
const SECRET_KEY = "bestsecretkey"

type Claims struct {
	jwt.RegisteredClaims
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("in auth middleware")
		next.ServeHTTP(w, r)
	})
}
