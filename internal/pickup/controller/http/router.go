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

	rt.Use(middleware.JWTVerify(eh.cfg))
	pickup := rt.Group("/api/pickup")
	{
		pickup.POST("/orders/:code/pickup", eh.Pickup)
		//pickup.POST("/orders/:code/refund")
		//pickup.POST("/orders/:code/receive")
		pickup.GET("/orders/:code", eh.GetOrderByCode)
		//pickup.POST("/orders/", eh.CreateOrder)
		pickup.GET("/orders/", eh.GetOrders)

		pickup.GET("/:user_id/pickup_orders", eh.GetPickupOrders)
		pickup.GET("/:user_id/pickup_orders/:pickup_order_id", eh.GetPickupOrderByID)

		pickup.GET("/customers/:customer_id", eh.GetCustomerByID)
		pickup.GET("/companies/:company_id", eh.GetCompanyByID)
	}

	return rt
}
