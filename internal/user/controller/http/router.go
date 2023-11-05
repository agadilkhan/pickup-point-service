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
		logger,
	}
}

func (r *router) GetHandler(eh *EndpointHandler) http.Handler {
	rt := gin.Default()

	user := rt.Group("/api/user/v1")
	{
		user.POST("/user/:login", eh.GetUserByLogin)
	}

	return rt
}
