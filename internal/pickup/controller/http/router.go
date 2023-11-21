package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/pickup/controller/http/middleware"
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

	rt.Use(middleware.JWTVerify(eh.cfg))
	pickup := rt.Group("/api/pickup")
	{
		pickup.POST("/orders/:code/pickup", eh.Pickup)
		pickup.POST("/orders/:code/refund")
		pickup.POST("/orders/:code/receive")
		pickup.GET("/orders/:code", eh.GetOrderByCode)
		pickup.POST("/orders/create", eh.CreateOrder)

		pickup.GET("/:user_id/pickup_orders")
		pickup.GET("/:user_id/pickup_orders/:id")
		pickup.GET("/:user_id/info")

		pickup.GET("/customers/:id", eh.GetCustomerByID)
		pickup.GET("/companies/:id", eh.GetCompanyByID)
	}

	return rt
}
