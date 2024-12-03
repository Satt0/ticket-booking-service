package shared

import (
	"http-server/internal/shared/database"
	"http-server/internal/shared/env"
	jwtutils "http-server/internal/shared/jwt"
	"http-server/internal/shared/kafka"
	"http-server/internal/shared/logger"
	"http-server/internal/shared/middleware"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type SharedDeps struct {
	Env         *env.Env
	Logger      *logger.Logger
	JwtUtils    *jwtutils.JwtUtils
	Middlewares *middleware.MiddleWare
	KafkaClient *kafka.KafkaClient
	DB          *gorm.DB
}

func NewSharedDeps(log *logger.Logger, db *gorm.DB, env *env.Env, jwtUtils *jwtutils.JwtUtils, middlewares *middleware.MiddleWare, kafkaClient *kafka.KafkaClient) *SharedDeps {
	return &SharedDeps{
		Env:         env,
		Logger:      log,
		JwtUtils:    jwtUtils,
		Middlewares: middlewares,
		KafkaClient: kafkaClient,
		DB:          db,
	}
}

var SharedModuleFx = fx.Options(
	fx.Provide(env.NewEnv),
	fx.Provide(logger.NewLogger),
	fx.Provide(jwtutils.NewJwtUtils),
	fx.Provide(NewSharedDeps),
	fx.Provide(middleware.NewMiddleWare),
	fx.Provide(database.NewDatabaseConnection),
	kafka.KafkaModuleFx,
)
