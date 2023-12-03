package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/agadilkhan/pickup-point-service/internal/auth/auth"
	"github.com/gin-gonic/gin"
)

func (h *EndpointHandler) GetUsers(ctx *gin.Context) {
	users, err := h.authService.GetUsers(ctx)
	if err != nil {
		h.logger.Errorf("failed to GetUsers err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *EndpointHandler) GetUserByID(ctx *gin.Context) {
	param := ctx.Param("user_id")

	userID, err := strconv.Atoi(param)
	if err != nil {
		h.logger.Errorf("failed to convert user_id to int: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	user, err := h.authService.GetUserByID(ctx, userID)
	if err != nil {
		h.logger.Errorf("failed to GetUserByID err: %v", err)
		ctx.Status(http.StatusNotFound)

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

	user, err := h.authService.UpdateUser(ctx, auth.UpdateUserRequest{
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

	respID, err := h.authService.DeleteUser(ctx, userID)
	if err != nil {
		h.logger.Errorf("failed to DeleteUser err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("user with id %d: deleted", respID))
}
