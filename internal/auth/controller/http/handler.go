package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/auth/auth"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"go.uber.org/zap"
)

type EndpointHandler struct {
	authService auth.UseCase
	logger      *zap.SugaredLogger
	cfg         *config.Config
}

func NewEndpointHandler(
	authService auth.UseCase,
	logger *zap.SugaredLogger,
	cfg *config.Config,
) *EndpointHandler {
	return &EndpointHandler{
		authService: authService,
		logger:      logger,
		cfg:         cfg,
	}
}
