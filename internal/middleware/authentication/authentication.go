package authentication

import (
	"net/http"
	"strings"
	"time"

	"github.com/DEHbNO4b/practicum_project/internal/authorization"
	"github.com/DEHbNO4b/practicum_project/internal/logger"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("in auth middleware")
		jwt := r.Header.Get("Authorization")
		jwt, found := strings.CutPrefix(jwt, "Bearer ")
		if !found || jwt == "" {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		claims := authorization.GetClaims(jwt)
		if claims.ExpiresAt.Before(time.Now()) {
			http.Error(w, "not allowed", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
