package middleware

import (
	"time"
	"weekly-task-3/configs"

	"github.com/golang-jwt/jwt"
)

func GetToken(id uint, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.Token))
}
