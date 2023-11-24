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

	auth := rt.Group("/api/auth/")
	{
		auth.POST("/register", eh.Register)
		auth.POST("/:email/user-confirm", eh.ConfirmUser)
		auth.POST("/login", eh.Login)
		auth.POST("/renew_token", eh.RenewToken)
	}

	return rt
}
