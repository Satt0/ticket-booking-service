package middleware

import (
	"http-server/internal/shared/env"
	jwtutils "http-server/internal/shared/jwt"
	"http-server/internal/shared/logger"
)

type MiddleWare struct {
	logger   *logger.Logger
	env      *env.Env
	jwtUtils *jwtutils.JwtUtils
}

func NewMiddleWare(logger *logger.Logger, env *env.Env, jwtUtils *jwtutils.JwtUtils) *MiddleWare {
	return &MiddleWare{logger: logger, env: env, jwtUtils: jwtUtils}
}
