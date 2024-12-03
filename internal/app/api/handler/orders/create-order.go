package ordershandler

import (
	users_dto "http-server/internal/app/api/handler/orders/dto"
	res_format "http-server/internal/shared/res-format"
	"http-server/internal/shared/utils"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) HandleCreateOrder(c *gin.Context) {
	dto, ok := utils.ValidateBodyDto[users_dto.CreateOrderReqDto](c)
	if !ok {
		return
	}
	res, err := uh.userServcie.CreatePendingOrder(dto)
	if err != nil {
		res_format.FormatResponse400(c, err.Error())
		return
	}
	res_format.FormatResponse200(c, res, nil)
}
