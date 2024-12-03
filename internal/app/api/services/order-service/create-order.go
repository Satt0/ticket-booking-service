package orderservice

import (
	users_dto "http-server/internal/app/api/handler/orders/dto"
	"http-server/internal/shared/database/entities"
)

func (s *OrderHistoryService) CreatePendingOrder(dto users_dto.CreateOrderReqDto) (*entities.Order, error) {
	// get event & ticket data
	return s.orderRepo.SaveOrderAndOutbox(dto)
}
