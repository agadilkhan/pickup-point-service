package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/user/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type EndpointHandler struct {
	userService user.UseCase
	logger      *zap.SugaredLogger
}

func NewEndpointHandler(
	userService user.UseCase,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		userService,
		logger,
	}
}

func (eh *EndpointHandler) GetUserByLogin(ctx *gin.Context) {
	login := ctx.Param("login")

	u, err := eh.userService.GetUserByLogin(ctx, login)
	if err != nil {
		eh.logger.Errorf("GetUserByLogin err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	response := struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Login     string `json:"login"`
		Password  string `json:"password"`
	}{
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.Phone,
		u.Login,
		u.Password,
	}

	ctx.JSON(http.StatusOK, response)
}

func (eh *EndpointHandler) CreateUser(ctx *gin.Context) {
	var request user.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		eh.logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	userID, err := eh.userService.CreateUser(ctx, request)
	if err != nil {
		eh.logger.Errorf("CreateUser request err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, userID)
}
