package auth_services

import (
	auth_dto "http-server/internal/app/api/handler/auth/dto"
	"http-server/internal/shared/database/entities"

	"github.com/gin-gonic/gin"
)

type UserAccessToken struct {
	Token string `json:"token"`
	Exp   uint   `json:"exp"`
}

func (as *AuthService) LoginUser(dto auth_dto.LoginDto) (interface{}, string) {
	user, err := as.userRepo.FindByEmail(dto.Email)
	if err != nil {
		return &entities.Users{}, "user not found"
	}
	if user.Password != dto.Password {
		return &entities.Users{}, "password not match"
	}
	token, exp, err := as.jwtUtils.EncryptPayload(user)
	if err != nil {
		return &entities.Users{}, "cannot generate token"
	}
	return gin.H{
		"user":  user,
		"token": &UserAccessToken{Token: token, Exp: exp},
	}, ""
}
