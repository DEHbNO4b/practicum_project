package authorisation

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TOKEN_EXP = time.Hour * 3
const SECRET_KEY = "bestsecretkey"

type Claims struct {
	jwt.RegisteredClaims
}
