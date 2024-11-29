package jwtutils

import (
	"github.com/golang-jwt/jwt"
)

func (jwtutils *JwtUtils) DecryptPayload(unVerifiedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(unVerifiedToken, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(jwtutils.env.JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
