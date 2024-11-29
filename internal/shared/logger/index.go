package logger

import (
	"http-server/internal/shared/env"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Out *logrus.Logger
}

func NewLogger(env *env.Env) *Logger {
	logger := logrus.New()
	// logger.SetLevel(logrus)
	return &Logger{Out: logger}
}
