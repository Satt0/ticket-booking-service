package auth_handler

import (
	"fmt"
	auth_dto "http-server/internal/app/api/handler/auth/dto"
	res_format "http-server/internal/shared/res-format"

	"github.com/gin-gonic/gin"
)

func (authHandler *AuthHandler) HandleSignUp(c *gin.Context) {
	var body *auth_dto.CreateUserDto
	err := c.ShouldBindBodyWithJSON(&body)
	if err != nil {
		fmt.Println(err)
		res_format.FormatResponse400(c, err.Error())
		return
	}
	newUser, errCode := authHandler.authService.SignUpUser(body)
	if errCode != "" {
		res_format.FormatResponse400(c, errCode)
		return
	}
	res_format.FormatResponse200(c, newUser, nil)
}
