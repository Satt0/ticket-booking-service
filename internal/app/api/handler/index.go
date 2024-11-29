package handler

import (
	auth_handler "http-server/internal/app/api/handler/auth"
	users_handler "http-server/internal/app/api/handler/users"

	"go.uber.org/fx"
)

var Handlers = fx.Options(fx.Provide(users_handler.NewUserHandler), fx.Provide(auth_handler.NewAuthHandler))
