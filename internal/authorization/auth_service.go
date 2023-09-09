package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func ParseTokenString(tokenString string) (MyJWTClaims, error) {
	claims := &MyJWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return MyJWTClaims{}, errors.New("signature invalid")
		} else {
			return MyJWTClaims{}, errors.New("can't parse token string")
		}
	}

	if !token.Valid {
		return MyJWTClaims{}, errors.New("token invalid")
	}

	return *claims, nil
}

func NewTokenStringWithClaims(claims MyJWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	return ss, err
}
