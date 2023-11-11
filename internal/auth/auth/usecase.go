package auth

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/auth/controller/http/dto"
)

type UseCase interface {
	GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JWTUserToken, error)
	RenewToken(ctx context.Context, refreshToken string) (*JWTUserToken, error)
	Register(ctx context.Context, request dto.CreateUserRequest) (int, error)
	//ValidateToken(tokenString string) (jwt.MapClaims, error)
}
