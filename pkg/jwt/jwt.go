package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Claims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(userId, email, jwtKey string, expirationTime time.Duration) (string, error) {
	var secret = jwtKey
	if secret == "" {
		secret = os.Getenv("JWT_SECRET")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			UserId:           userId,
			Email:            email,
			RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(expirationTime)}},
		})
	token, err := t.SignedString([]byte(jwtKey))
	if err != nil {
		return "", nil
	}
	return token, nil
}

// ParseToken 解析token
func ParseToken(tokenString, jwtKey string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	var secret = jwtKey
	if secret == "" {
		secret = os.Getenv("JWT_SECRET")
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	return token, claims, err
}
