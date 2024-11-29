package auth_services

import (
	"http-server/internal/shared"
	"http-server/internal/shared/database/repository"
	jwtutils "http-server/internal/shared/jwt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	deps     *shared.SharedDeps
	jwtUtils *jwtutils.JwtUtils
}

func NewAuthService(userRepo *repository.UserRepository, jwtUtils *jwtutils.JwtUtils, deps *shared.SharedDeps) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtUtils: jwtUtils,
		deps:     deps,
	}
}
