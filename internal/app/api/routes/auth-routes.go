package routes

import (
	auth_handler "http-server/internal/app/api/handler/auth"

	"github.com/gin-gonic/gin"
)

type AuthRouting struct {
	r       *gin.Engine
	handler *auth_handler.AuthHandler
}

func (ur *AuthRouting) SetUp() {
	g2 := ur.r.Group("/auth")
	{
		g2.POST("/login", ur.handler.HandleLoginUser)
		g2.POST("/signup", ur.handler.HandleSignUp)

	}
}
func NewAuthRouting(r *gin.Engine, auth_handler *auth_handler.AuthHandler) *AuthRouting {
	return &AuthRouting{
		r:       r,
		handler: auth_handler,
	}
}
