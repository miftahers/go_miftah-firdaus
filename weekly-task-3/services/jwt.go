package services

import (
	"errors"
	"strings"
	"weekly-task-3/configs"
	"weekly-task-3/dto"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func DecodeJWT(ctx echo.Context) (dto.Token, error) {
	var t dto.Token

	auth := ctx.Request().Header.Get("Authorization")
	if auth == "" {
		return dto.Token{}, errors.New("authorization header not found")
	}
	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]

	token, err := jwt.ParseWithClaims(auth, &dto.Token{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.Token), nil
	})
	if err != nil {
		return dto.Token{}, errors.New("token is wrong or expired")
	}
	if claims, ok := token.Claims.(*dto.Token); ok && token.Valid {
		t.UserID = claims.UserID
		t.Username = claims.Username
	}
	return t, nil
}
