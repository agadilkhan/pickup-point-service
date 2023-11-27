package http

import (
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/admin/admin"
	"github.com/agadilkhan/pickup-point-service/internal/admin/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type EndpointHandler struct {
	service admin.UseCase
	logger  *zap.SugaredLogger
	cfg     *config.Config
}

func NewEndpointHandler(
	service admin.UseCase,
	logger *zap.SugaredLogger,
	cfg *config.Config,
) *EndpointHandler {
	return &EndpointHandler{
		service: service,
		logger:  logger,
		cfg:     cfg,
	}
}

func (h *EndpointHandler) GetUsers(ctx *gin.Context) {

}

func (h *EndpointHandler) GetUserByID(ctx *gin.Context) {
	param := ctx.Param("user_id")

	userID, err := strconv.Atoi(param)
	if err != nil {
		h.logger.Errorf("failed to convert user_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	user, err := h.service.GetUserByID(ctx, userID)
	if err != nil {
		h.logger.Errorf("failed to GetUserByID err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *EndpointHandler) UpdateUser(ctx *gin.Context) {
	param := ctx.Param("user_id")

	userID, err := strconv.Atoi(param)
	if err != nil {
		h.logger.Errorf("failed to convert user_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	request := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Login     string `json:"login"`
		Password  string `json:"password"`
	}{}

	if err = ctx.ShouldBindJSON(&request); err != nil {
		h.logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	user, err := h.service.UpdateUser(ctx, admin.UpdateUserRequest{
		ID:        userID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Login:     request.Login,
		Password:  request.Password,
	})
	if err != nil {
		h.logger.Errorf("failed to UpdateUser err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *EndpointHandler) DeleteUser(ctx *gin.Context) {
	param := ctx.Param("user_id")

	userID, err := strconv.Atoi(param)
	if err != nil {
		h.logger.Errorf("failed to convert user_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	respID, err := h.service.DeleteUser(ctx, userID)
	if err != nil {
		h.logger.Errorf("failed to DeleteUser err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("user with id %d: deleted", respID))
}
