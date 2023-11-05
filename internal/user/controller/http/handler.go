package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/user/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EndpointHandler struct {
	userService user.UseCase
	logger      *zap.SugaredLogger
}

func NewEndpointHandler(
	userService user.UseCase,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		userService,
		logger,
	}
}

func (eh *EndpointHandler) GetUserByLogin(ctx *gin.Context) {

}
