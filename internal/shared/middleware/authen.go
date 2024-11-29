package middleware

import (
	res_format "http-server/internal/shared/res-format"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (m *MiddleWare) CreateAuthUserMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		defer func() {
			err := recover()
			if err != nil {
				m.logger.Out.Warn(err)
				res_format.FormatResponse401(c)
				c.Abort()
			}
		}()
		if authHeader == "" {
			panic("empty Authen header")
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) < 2 || parts[0] != "Bearer" {
			panic("mal-format Authen header")
		}
		unVerifiedToken := parts[1]
		token, err := m.jwtUtils.DecryptPayload(unVerifiedToken)
		if err != nil || !token.Valid {
			panic("authMiddleware:33: " + err.Error())
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Set user information to context
			c.Set("user", claims["sub"]) // Assuming "sub" contains user ID
		} else {
			panic("cannot map jwt data to user info")
		}
	}

}
