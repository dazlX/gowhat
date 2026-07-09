package transport

import (
	"gowhat/internal/repository"

	"go.uber.org/zap"
)

type HandlerConfig struct {
	*zap.Logger
	Repo *repository.UserRepository
}

func NewHandlerConfig(log *zap.Logger, repo *repository.UserRepository) *HandlerConfig {
	return &HandlerConfig{
		Logger: log,
		Repo:   repo,
	}
}
