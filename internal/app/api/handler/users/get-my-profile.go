package users_handler

import (
	res_format "http-server/internal/shared/res-format"
	"http-server/internal/shared/utils"

	"github.com/gin-gonic/gin"
)

func (uh *UserHandler) HandleGetMyProfile(c *gin.Context) {
	userId, err := utils.GetUserIdFromContext(c)
	if err != nil {
		res_format.FormatResponse400(c, err.Error())
		return
	}
	profile, errCode := uh.userServcie.GetUserById(userId)

	if errCode != "" {
		res_format.FormatResponse404(c, errCode)
		return
	}
	res_format.FormatResponse200(c, profile, nil)
}
