package http

import (
	"github.com/agadilkhan/pickup-point-service/internal/auth/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type EndpointHandler struct {
	authService auth.UseCase
	logger      *zap.SugaredLogger
}

func NewEndpointHandler(
	authService auth.UseCase,
	logger *zap.SugaredLogger,
) *EndpointHandler {
	return &EndpointHandler{
		authService: authService,
		logger:      logger,
	}
}

func (h *EndpointHandler) Login(ctx *gin.Context) {
	logger := h.logger.With(
		zap.String("endpoint", "login"),
		zap.String("params", ctx.FullPath()),
	)

	request := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("failed to unmarshal body err: %v", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	tokenRequest := auth.GenerateTokenRequest{
		Login:    request.Login,
		Password: request.Password,
	}

	userToken, err := h.authService.GenerateToken(ctx, tokenRequest)
	if err != nil {
		logger.Errorf("failed to GenerateToken err: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	response := struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *EndpointHandler) Register(ctx *gin.Context) {
	logger := h.logger.With(
		zap.String("endpoint", "register"),
		zap.String("params", ctx.FullPath()),
	)

	request := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	createUserRequest := auth.CreateUserRequest{
		Login:    request.Login,
		Password: request.Password,
	}

	userID, err := h.authService.Register(ctx, createUserRequest)
	if err != nil {
		logger.Errorf("Register request err: %v", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, userID)
}

func (h *EndpointHandler) RenewToken(ctx *gin.Context) {

}
