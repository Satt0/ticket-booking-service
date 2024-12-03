package handler

import (
	ordershandler "http-server/internal/app/api/handler/orders"

	"go.uber.org/fx"
)

var Handlers = fx.Options(fx.Provide(ordershandler.NewUserHandler))
