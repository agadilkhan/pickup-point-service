package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/auth/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EndpointHandler struct {
	authService auth.UseCase
	logger      *zap.SugaredLogger
}

func NewEndpointHandler(
	authService auth.UseCase,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *EndpointHandler) Login(ctx *gin.Context) {

}
