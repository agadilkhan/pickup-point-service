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
		api.POST("/orders", eh.CreateOrder)
		api.GET("/orders", eh.GetOrders)
		api.GET("/orders/:order_code", eh.GetOrderByCode)
		api.PUT("/orders/:order_code")
		api.DELETE("/orders/:order_code", eh.DeleteOrder)

		api.POST("/orders/:order_code/pickup", eh.PickupOrder)

		api.POST("/orders/:order_code/receive", eh.ReceiveOrder)
		api.POST("/orders/:order_code/items/:product_id/receive", eh.ReceiveItem)

		api.POST("/orders/:order_code/cancel", eh.CancelOrder)

		api.POST("/orders/:order_code/refund", eh.RefundOrder)
		api.POST("/orders/:order_code/items/:product_id/refund", eh.RefundItem)

		api.GET("/:user_id/transactions", eh.GetTransactions)
	}
	return rt
}
