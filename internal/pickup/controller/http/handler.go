package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/pickup"
	"go.uber.org/zap"
)

type EndpointHandler struct {
	service *pickup.Service
	logger  *zap.SugaredLogger
}

func NewEndpointHandler(
	service *pickup.Service,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		service: service,
		logger:  logger,
	}
}
