package users_handler

import (
	"fmt"
	users_dto "http-server/internal/app/api/handler/users/dto"
	user_services "http-server/internal/app/api/services/user-services"
	"http-server/internal/shared/pagination"
	res_format "http-server/internal/shared/res-format"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userServcie *user_services.UserService
}

func NewUserHandler(userServcie *user_services.UserService) *UserHandler {
	return &UserHandler{
		userServcie: userServcie,
	}
}

func (h *UserHandler) HandleGetAllUser(c *gin.Context) {
	// dto validation
	var param users_dto.GetManyUserReqDto
	err := c.ShouldBindQuery(&param)

	if err != nil {
		fmt.Println(err)
		res_format.FormatResponse400(c, err.Error())
		return
	}

	users := h.userServcie.GetAllUser()
	pg := pagination.GetPaginationRespone(101, pagination.PaginationRequestDto{Limit: param.Limit, Page: param.Page})

	res_format.FormatResponse200(c, users, pg)
}
func (h *UserHandler) HandleGetUserById(c *gin.Context) {
	var param users_dto.GetUserByIdReqDto
	err := c.ShouldBindQuery(&param)
	if err != nil {
		fmt.Println(err)
		res_format.FormatResponse400(c, err.Error())
		return
	}
	user, errCode := h.userServcie.GetUserById(param.Id)
	if errCode != "" {
		fmt.Println(errCode)
		res_format.FormatResponse404(c, errCode)
		return
	}
	res_format.FormatResponse200(c, user, nil)
}
