package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http/middleware"
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

func (r *router) GetHandler(h *EndpointHandler) http.Handler {
	rt := gin.Default()

	admin := rt.Group("/api/admin/v1")
	{
		admin.Use(middleware.JWTVerify(h.cfg, h.logger))
		admin.GET("/users/", h.GetUsers)
		admin.GET("/users/:user_id", h.GetUserByID)
		admin.PUT("/users/:user_id", h.UpdateUser)
		admin.DELETE("/users/:user_id", h.DeleteUser)
	}

	return rt
}
