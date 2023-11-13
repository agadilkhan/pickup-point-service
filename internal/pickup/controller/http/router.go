package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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

	order := rt.Group("/api/pickup")
	{
		order.GET("/:id")
	}

	return rt
}
