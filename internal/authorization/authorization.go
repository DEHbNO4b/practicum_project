package authorization

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TOKEN_EXP = time.Minute * 1
const SECRET_KEY = "bestsecretkey"

type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

func BuildJWTString() (string, error) {
	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
		// собственное утверждение
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}

func GetClaims(tokenString string) Claims {
	// создаём экземпляр структуры с утверждениями
	claims := Claims{}
	// парсим из строки токена tokenString в структуру claims
	jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	// возвращаем ID пользователя в читаемом виде
	return claims
}
