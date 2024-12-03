package users_dto

import "http-server/internal/shared/pagination"

type GetManyUserReqDto struct {
	pagination.PaginationRequestDto
}

type GetUserByIdReqDto struct {
	Id uint64 `form:"id" binding:"required"`
}

type OrderTicketOptionReqDto struct {
	OptionId uint64 `json:"optionId" binding:"required" validate:"min=1"`
	Amount   uint64 `json:"amount" binding:"required" validate:"min=1"`
}
type CreateOrderReqDto struct {
	EventId uint64                    `json:"eventId" binding:"required" validate:"min=1"`
	Cart    []OrderTicketOptionReqDto `validate:"required,dive,required"`
}
