package pagination

import "math"

type PaginationResponseDto struct {
	Limit       uint64 `json:"limit"`
	Page        uint64 `json:"page"`
	TotalRecord uint64 `json:"total_record"`
	TotalPage   uint64 `json:"total_page"`
}
type PaginationRequestDto struct {
	Limit uint64 `form:"limit" binding:"required"`
	Page  uint64 `form:"page" binding:"required"`
}

func GetPaginationRespone(count uint64, p PaginationRequestDto) *PaginationResponseDto {
	return &PaginationResponseDto{
		Limit:       p.Limit,
		Page:        p.Page,
		TotalRecord: count,
		TotalPage:   uint64(math.Ceil(float64(count) / float64(p.Limit))),
	}
}
