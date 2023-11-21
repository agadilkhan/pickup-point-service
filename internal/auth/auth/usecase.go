package auth

import (
	"context"
)

type UseCase interface {
	GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JWTUserToken, error)
	RenewToken(ctx context.Context, refreshToken string) (*JWTUserToken, error)
	Register(ctx context.Context, request CreateUserRequest) (int, error)
	ConfirmUser(cxt context.Context, code ConfirmUserRequest) error
}
