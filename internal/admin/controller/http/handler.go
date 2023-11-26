package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/admin/admin"
	"go.uber.org/zap"
)

type EndpointHandler struct {
	service admin.UseCase
	logger  *zap.SugaredLogger
}

func NewEndpointHandler(
	service admin.UseCase,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		service: service,
		logger:  logger,
	}
}
