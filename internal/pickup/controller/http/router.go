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

func (r *router) GetHandler(eh *EndpointHandler) http.Handler {
	rt := gin.Default()

	rt.Use(middleware.JWTVerify(eh.cfg))
	pickup := rt.Group("/api/pickup")
	{
		orders := pickup.Group("/orders")
		{
			orders.GET("/", eh.GetOrders)
			orders.POST("/", eh.CreateOrder)
			orders.GET("/:code", eh.GetOrderByCode)
			//orders.PUT("/:code")
			//orders.DELETE("/:code")
			orders.POST("/:code/pickup", eh.Pickup)
			orders.POST("/:code/receive", eh.ReceiveOrder)
		}
		pickupOrders := pickup.Group("/:user_id/pickup_orders")
		{
			pickupOrders.GET("/", eh.GetPickupOrders)
			pickupOrders.GET("/:pickup_order_id", eh.GetPickupOrderByID)
		}
		pickupPoints := pickup.Group("/pickup_points")
		{
			pickupPoints.GET("/", eh.GetAllPickupPoints)
			pickupPoints.GET("/:pickup_point_id", eh.GetPickupPointByID)
		}
		customers := pickup.Group("/customers")
		{
			customers.GET("/", eh.GetAllCustomers)
			customers.GET("/:customer_id", eh.GetCustomerByID)
		}
		companies := pickup.Group("/companies")
		{
			companies.GET("/", eh.GetAllCompanies)
			companies.GET("/company_id", eh.GetCompanyByID)
		}
		products := pickup.Group("/products")
		{
			products.GET("/", eh.GetAllProducts)
			products.GET("/:product_id", eh.GetProductByID)
		}
		warehouses := pickup.Group("/warehouses")
		{
			warehouses.GET("/", eh.GetAllWarehouses)
			warehouses.GET("/:warehouse_id", eh.GetWarehouseByID)
		}
	}

	return rt
}
