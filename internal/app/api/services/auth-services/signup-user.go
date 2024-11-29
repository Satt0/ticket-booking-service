package auth_services

import (
	"fmt"
	auth_dto "http-server/internal/app/api/handler/auth/dto"
	"http-server/internal/shared/database/entities"

	"github.com/gin-gonic/gin"
)

func (as *AuthService) SignUpUser(dto *auth_dto.CreateUserDto) (interface{}, string) {
	_, err := as.userRepo.FindByEmail(dto.Email)
	if err == nil {
		return &entities.Users{}, "user duplicated"
	}

	newUser, err := as.userRepo.Create(entities.Users{
		Email:    dto.Email,
		Password: dto.Password,
		Name:     dto.Name,
		Age:      "18",
		Gender:   "Male",
		Balance:  "0",
	})
	if err != nil {
		return &entities.Users{}, err.Error()
	}
	token, exp, err := as.jwtUtils.EncryptPayload(newUser)
	if err != nil {
		return &entities.Users{}, "cannot generate token"
	}
	fmt.Printf("user: %v\n", newUser)
	return gin.H{
		"user":  newUser,
		"token": &UserAccessToken{Token: token, Exp: exp},
	}, ""
}
