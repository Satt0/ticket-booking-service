package repository

import "go.uber.org/fx"

var Repositories = fx.Options(fx.Provide(NewUserRepository), fx.Provide(NewTicketOptionRepository))
