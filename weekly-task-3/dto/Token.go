package dto

import "github.com/golang-jwt/jwt"

type Token struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
