package auth

import (
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"github.com/agadilkhan/pickup-point-service/internal/auth/repository"
	"github.com/agadilkhan/pickup-point-service/internal/auth/transport"
	"github.com/agadilkhan/pickup-point-service/internal/kafka"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	repo                     repository.Repository
	jwtSecretKey             string
	passwordSecretKey        string
	userGrpcTransport        *transport.UserGrpcTransport
	userVerificationProducer *kafka.Producer
	redisCli                 *redis.Client
}

func NewAuthService(
	repo repository.Repository,
	authConfig config.Auth,
	userGrpcTransport *transport.UserGrpcTransport,
	userVerificationProducer *kafka.Producer,
	redisCli *redis.Client,
) UseCase {
	return &Service{
		repo:                     repo,
		jwtSecretKey:             authConfig.JWTSecretKey,
		passwordSecretKey:        authConfig.PasswordSecretKey,
		userGrpcTransport:        userGrpcTransport,
		userVerificationProducer: userVerificationProducer,
		redisCli:                 redisCli,
	}
}
