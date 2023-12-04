package http

import (
	"net/http"

	"github.com/agadilkhan/pickup-point-service/internal/auth/controller/http/middleware"

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

	// swagger
	rt.GET("/swagger/*any", gin.WrapH(h.Swagger()))

	api := rt.Group("/api/auth/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/register", h.Register)
			user.POST("/login", h.Login)
			user.POST("/:email/confirm-user", h.ConfirmUser)
			user.POST("/renew-token", h.RenewToken)
		}

		admin := api.Group("/admin")
		{
			admin.Use(middleware.AdminMiddleware(h.cfg, h.logger))
			admin.GET("/users", h.GetUsers)
			admin.GET("/users/:user_id", h.GetUserByID)
			admin.PUT("/users/:user_id", h.UpdateUser)
			admin.DELETE("/users/:user_id", h.DeleteUser)
		}
	}

	return rt
}
