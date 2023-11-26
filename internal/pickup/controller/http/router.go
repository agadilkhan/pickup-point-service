package http

import (
	"net/http"

	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http/middleware"
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

	rt.Use(middleware.JWTVerify(eh.cfg, eh.logger))
	api := rt.Group("/api/pickup/v1")
	{
		eh.initOrderRoutes(api)
		eh.initPickupRoutes(api)
		eh.initCustomerRoutes(api)
		eh.initProductRoutes(api)
		eh.initCompanyRoutes(api)
		eh.initWarehouseRoutes(api)
	}
	return rt
}
