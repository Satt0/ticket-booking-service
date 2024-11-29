package services

import (
	auth_services "http-server/internal/app/api/services/auth-services"
	user_services "http-server/internal/app/api/services/user-services"

	"go.uber.org/fx"
)

var Services = fx.Options(fx.Provide(user_services.NewUserService), fx.Provide(auth_services.NewAuthService))
