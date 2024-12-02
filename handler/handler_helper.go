package handler

import (
	"github.com/dargasht/gocrud/database/repo"
	"go.uber.org/zap"
)

type HandlerConfig struct {
	DB     *repo.Queries
	Logger *zap.Logger
}

func NewHandlerConfig(db *repo.Queries, logger *zap.Logger) *HandlerConfig {
	return &HandlerConfig{
		DB:     db,
		Logger: logger,
	}
}
