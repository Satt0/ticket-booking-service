package utils

import (
	res_format "http-server/internal/shared/res-format"

	"github.com/gin-gonic/gin"
)

func ValidateBodyDto[k interface{}](c *gin.Context) (k, bool) {
	var param k
	err := c.ShouldBindBodyWithJSON(&param)

	if err != nil {
		res_format.FormatResponse400(c, err.Error())
		c.Abort()
		return param, false
	}
	return param, true
}
