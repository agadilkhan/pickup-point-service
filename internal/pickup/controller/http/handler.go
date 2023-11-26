package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/config"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"go.uber.org/zap"
)

type EndpointHandler struct {
	service pickup.UseCase
	logger  *zap.SugaredLogger
	cfg     *config.Config
}

func NewEndpointHandler(
	service pickup.UseCase,
	logger *zap.SugaredLogger,
	cfg *config.Config,
) *EndpointHandler {
	return &EndpointHandler{
		service: service,
		logger:  logger,
		cfg:     cfg,
	}
}
