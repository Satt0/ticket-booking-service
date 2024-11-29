package jwtutils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func (jwtUtils *JwtUtils) EncryptPayload(data interface{}) (string, uint, error) {
	exp := uint(time.Now().Add(time.Hour).Unix())
	claims := jwt.MapClaims{
		"sub": data,
		"exp": exp, // Expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtUtils.env.JWT_SECRET))

	if err != nil {
		return "", 0, err
	}

	return tokenString, exp, nil
}
