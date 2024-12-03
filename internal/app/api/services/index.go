package services

import (
	orderservice "http-server/internal/app/api/services/order-service"

	"go.uber.org/fx"
)

var Services = fx.Options(fx.Provide(orderservice.NewOrderHistoryService))
