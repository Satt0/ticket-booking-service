package orderservice

import "http-server/internal/shared/database/entities"

func (s *OrderHistoryService) GetUserOrderHistory() (*entities.Order, error) {
	return s.orderRepo.FindByID(1)
	// return s.orderRepo.FindMany()
}
