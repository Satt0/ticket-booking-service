package routes

import (
	users_handler "http-server/internal/app/api/handler/users"
	"http-server/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouting struct {
	r          *gin.Engine
	handlers   *users_handler.UserHandler
	middleware *middleware.MiddleWare
}

func (ur *UserRouting) SetUp() {
	group := ur.r.Group("/users")
	{
		publicRoutes := group.Group("")
		{
			publicRoutes.GET("", ur.handlers.HandleGetAllUser)
			publicRoutes.GET("/by-id", ur.handlers.HandleGetUserById)
		}
		privateRoutes := group.Group("")
		{
			privateRoutes.Use(ur.middleware.CreateAuthUserMiddleWare())
			privateRoutes.GET("/my-profile", ur.handlers.HandleGetMyProfile)
		}

	}

}
func NewUserRouting(r *gin.Engine, uh *users_handler.UserHandler, m *middleware.MiddleWare) *UserRouting {
	return &UserRouting{
		r:          r,
		handlers:   uh,
		middleware: m,
	}
}
