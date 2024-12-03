package routes

import (
	ordershandler "http-server/internal/app/api/handler/orders"
	"http-server/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouting struct {
	r          *gin.Engine
	handlers   *ordershandler.UserHandler
	middleware *middleware.MiddleWare
}

func (ur *UserRouting) SetUp() {
	group := ur.r.Group("orders")
	{
		group.POST("", ur.handlers.HandleCreateOrder)
		group.GET("/all", ur.handlers.HandleGetMyProfile)
	}

}
func NewUserRouting(r *gin.Engine, uh *ordershandler.UserHandler, m *middleware.MiddleWare) *UserRouting {
	return &UserRouting{
		r:          r,
		handlers:   uh,
		middleware: m,
	}
}
