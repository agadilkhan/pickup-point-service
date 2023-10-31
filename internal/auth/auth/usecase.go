package auth

import "context"

type UseCase interface {
	GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JWTUserToken, error)
}
