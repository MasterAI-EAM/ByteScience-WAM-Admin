package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct{}

func (j JWT) GetToken(secretKey string, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + seconds
	claims["iat"] = time.Now().Unix()
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
