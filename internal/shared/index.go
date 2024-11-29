package shared

import (
	"http-server/internal/shared/database"
	"http-server/internal/shared/env"
	jwtutils "http-server/internal/shared/jwt"
	"http-server/internal/shared/kafka"
	"http-server/internal/shared/logger"
	"http-server/internal/shared/middleware"

	"go.uber.org/fx"
)

type SharedDeps struct {
	Env         *env.Env
	Logger      *logger.Logger
	JwtUtils    *jwtutils.JwtUtils
	Middlewares *middleware.MiddleWare
	KafkaClient *kafka.KafkaClient
}

func NewSharedDeps(log *logger.Logger, env *env.Env, jwtUtils *jwtutils.JwtUtils, middlewares *middleware.MiddleWare, kafkaClient *kafka.KafkaClient) *SharedDeps {
	return &SharedDeps{
		Env:         env,
		Logger:      log,
		JwtUtils:    jwtUtils,
		Middlewares: middlewares,
		KafkaClient: kafkaClient,
	}
}

var SharedModuleFx = fx.Options(
	fx.Provide(env.NewEnv),
	fx.Provide(logger.NewLogger),
	fx.Provide(jwtutils.NewJwtUtils),
	fx.Provide(NewSharedDeps),
	fx.Provide(middleware.NewMiddleWare),
	database.DBConnection,
	kafka.KafkaModuleFx,
)
