package ordershandler

import orderservice "http-server/internal/app/api/services/order-service"

type UserHandler struct {
	userServcie *orderservice.OrderHistoryService
}

func NewUserHandler(userServcie *orderservice.OrderHistoryService) *UserHandler {
	return &UserHandler{
		userServcie: userServcie,
	}
}
