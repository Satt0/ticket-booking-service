package users_dto

import "http-server/internal/shared/pagination"

type GetManyUserReqDto struct {
	pagination.PaginationRequestDto
}

type GetUserByIdReqDto struct {
	Id uint64 `form:"id" binding:"required"`
}
