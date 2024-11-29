package auth_handler

import (
	auth_dto "http-server/internal/app/api/handler/auth/dto"
	auth_services "http-server/internal/app/api/services/auth-services"
	res_format "http-server/internal/shared/res-format"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *auth_services.AuthService
}

func NewAuthHandler(authService *auth_services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) HandleLoginUser(c *gin.Context) {

	param, ok := ValidateDto[auth_dto.LoginDto](c)

	if !ok {
		return
	}

	userLoginResp, errCode := h.authService.LoginUser(param)

	if errCode != "" {
		res_format.FormatResponse404(c, errCode)
		return
	}

	res_format.FormatResponse200(c, userLoginResp, nil)
}

func ValidateDto[k interface{}](c *gin.Context) (k, bool) {
	var param k
	err := c.ShouldBindBodyWithJSON(&param)

	if err != nil {
		res_format.FormatResponse400(c, err.Error())
		c.Abort()
		return param, false
	}
	return param, true
}
