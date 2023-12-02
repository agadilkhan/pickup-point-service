package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type router struct {
	logger *zap.SugaredLogger
}

func NewRouter(logger *zap.SugaredLogger) *router {
	return &router{
		logger: logger,
	}
}

func (r *router) GetHandler(eh *EndpointHandler) http.Handler {
	rt := gin.Default()

	api := rt.Group("/api/auth/v1")
	{
		api.POST("/register", eh.Register)
		api.POST("/:email/user-confirm", eh.ConfirmUser)
		api.POST("/login", eh.Login)
		api.POST("/renew_token", eh.RenewToken)
	}

	return rt
}
