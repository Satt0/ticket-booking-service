package jwtutils

import (
	"http-server/internal/shared/env"
	"http-server/internal/shared/logger"
)

type JwtUtils struct {
	log *logger.Logger
	env *env.Env
}

func NewJwtUtils(log *logger.Logger, env *env.Env) *JwtUtils {
	return &JwtUtils{
		log: log,
		env: env,
	}
}
