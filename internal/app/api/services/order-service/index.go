package orderservice

import "http-server/internal/shared/database/repository"

type OrderHistoryService struct {
	orderRepo *repository.OrderRepository
}

func NewOrderHistoryService(orderRepo *repository.OrderRepository) *OrderHistoryService {

	return &OrderHistoryService{
		orderRepo: orderRepo,
	}
}
