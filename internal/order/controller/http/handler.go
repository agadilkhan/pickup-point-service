package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/order/order"
	"go.uber.org/zap"
)

type EndpointHandler struct {
	service *order.Service
	logger  *zap.SugaredLogger
}

func NewEndpointHandler(
	service *order.Service,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		service: service,
		logger:  logger,
	}
}
