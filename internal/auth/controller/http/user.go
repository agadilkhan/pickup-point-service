package http

import (
	"fmt"
	"net/http"

	"github.com/agadilkhan/pickup-point-service/internal/auth/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// swagger:route POST /v1/user/login Login
//
//				Parameters:
//				 + name: GenerateToken
//	           in: body
//				type: GenerateToken
//
//				Consumes:
//				- application/json
//
//				Produces:
//				- application/json
//
//				Schemes: http, https
//
//
//				Responses:
//				  200: UserToken
//			   400:
//			   500:
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

//			swagger:route POST /v1/user/register Register
//
//			Parameters:
//			 + name: RegisterUser
//			   in: body
//	 		type: RegisterUser
//
//			Consumes:
//			- application/json
//
//			Produces:
//			- application/json
//
//			Schemes: http, https
//
//
//			Responses:
//			  201: ResponseCreated
//		   	  400:
//		   	  500:
func (h *EndpointHandler) Register(ctx *gin.Context) {
	logger := h.logger.With(
		zap.String("endpoint", "register"),
		zap.String("params", ctx.FullPath()),
	)

	request := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Login     string `json:"login"`
		Password  string `json:"password"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	createUserRequest := auth.CreateUserRequest{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Login:     request.Login,
		Password:  request.Password,
	}

	userID, err := h.authService.Register(ctx, createUserRequest)
	if err != nil {
		logger.Errorf("failed to Register err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusCreated, responseCreated{
		ID: userID,
	})
}

//			swagger:route POST /v1/user/renew-token RenewToken
//
//			Parameters:
//			+ name: RefreshToken
//			in: body
//	     		type: RefreshToken
//
//			Consumes:
//			- application/json
//
//			Produces:
//			- application/json
//
//			Schemes: http, https
//
//
//			Responses:
//			  200: UserToken
//		   	  400:
//		   	  500:
func (h *EndpointHandler) RenewToken(ctx *gin.Context) {
	logger := h.logger.With(
		zap.String("endpoint", "renew_token"),
		zap.String("params", ctx.FullPath()),
	)

	request := struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	jwtToken, err := h.authService.RenewToken(ctx, request.RefreshToken)
	if err != nil {
		logger.Errorf("failed to RenewToken err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	response := struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		jwtToken.Token,
		jwtToken.RefreshToken,
	}

	ctx.JSON(http.StatusOK, responseOK{
		response,
	})
}

// swagger:route POST /v1/user/{email}/confirm-user ConfirmUser
//
//		Consumes:
//		- application/json
//
//		Produces:
//		- application/json
//
//		Schemes: http, https
//
//		Parameters:
//		 + name: email
//		   in: path
//		+ name: ConfirmUserRequest
//			in: body
//			type: ConfirmUserRequest
//
//
//		Responses:
//	 200: ResponseMessage
//	 400:
//	 500:
func (h *EndpointHandler) ConfirmUser(ctx *gin.Context) {
	logger := h.logger.With(
		zap.String("endpoint", "renew_token"),
		zap.String("params", ctx.FullPath()),
	)

	email := ctx.Param("email")

	request := struct {
		Code string `json:"code"`
	}{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Errorf("failed to Unmarshal err: %v", err)
		ctx.Status(http.StatusBadRequest)

		return
	}

	userCode := auth.ConfirmUserRequest{
		Email: email,
		Code:  request.Code,
	}

	err := h.authService.ConfirmUser(ctx, userCode)
	if err != nil {
		logger.Errorf("failed to ConfirmUser err: %v", err)
		ctx.Status(http.StatusInternalServerError)

		return
	}

	ctx.JSON(http.StatusOK, responseMessage{
		Message: fmt.Sprintf("user with email %s: confirmed", email),
	})
}
