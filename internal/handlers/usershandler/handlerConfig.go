package usershandler

import (
	"gowhat/internal/service/users"

	"go.uber.org/zap"
)

type HandlerConfig struct {
	*zap.Logger
	Service users.Service
}

func NewHandlerConfig(log *zap.Logger, service *users.Service) *HandlerConfig {
	return &HandlerConfig{
		Logger:  log,
		Service: *service,
	}
}
