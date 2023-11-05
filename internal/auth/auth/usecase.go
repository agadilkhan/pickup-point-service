package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type UseCase interface {
	GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JWTUserToken, error)
	RenewToken(ctx context.Context, refreshToken string) (*JWTUserToken, error)
	Register(ctx context.Context) (int, error)
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}
