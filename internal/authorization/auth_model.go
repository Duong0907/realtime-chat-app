package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
