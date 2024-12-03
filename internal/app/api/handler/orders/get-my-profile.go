package ordershandler

import (
	res_format "http-server/internal/shared/res-format"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) HandleGetMyProfile(c *gin.Context) {
	res, err := uh.userServcie.GetUserOrderHistory()
	if err != nil {
		res_format.FormatResponse400(c, err.Error())
		return
	}
	res_format.FormatResponse200(c, res, nil)
}
