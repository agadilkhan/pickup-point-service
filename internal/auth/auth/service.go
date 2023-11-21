package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"github.com/agadilkhan/pickup-point-service/internal/auth/controller/consumer/dto"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
	"github.com/agadilkhan/pickup-point-service/internal/auth/repository"
	"github.com/agadilkhan/pickup-point-service/internal/auth/transport"
	"github.com/agadilkhan/pickup-point-service/internal/kafka"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
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

func (s *Service) GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JWTUserToken, error) {
	user, err := s.userGrpcTransport.GetUserByLogin(ctx, request.Login)

	generatedPassword := s.generatePassword(request.Password)
	if user.Password != generatedPassword {
		return nil, fmt.Errorf("password is wrong")
	}

	type MyCustomClaims struct {
		UserID int `json:"user_id"`
		RoleID int `json:"role_id"`
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		int(user.Id),
		int(user.RoleId),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	secretKey := []byte(s.jwtSecretKey)
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := claimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	rClaims := MyCustomClaims{
		int(user.Id),
		int(user.RoleId),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(40 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	rClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)

	refreshTokenString, err := rClaimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %w", err)
	}

	userToken := entity.UserToken{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
		UserID:       int(user.Id),
	}

	err = s.repo.CreateUserToken(ctx, userToken)
	if err != nil {
		return nil, fmt.Errorf("CreateUserToken err: %w", err)
	}

	jwtToken := &JWTUserToken{
		Token:        userToken.Token,
		RefreshToken: userToken.RefreshToken,
	}

	return jwtToken, nil
}

func (s *Service) RenewToken(ctx context.Context, refreshToken string) (*JWTUserToken, error) {
	claims, err := s.validateToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("ValidateToken err: %v", err)
	}

	userID, ok := claims["user_id"]
	if !ok {
		return nil, fmt.Errorf("user_id could not be parsed from JWT")
	}

	roleID, ok := claims["role_id"]
	if !ok {
		return nil, fmt.Errorf("role_id could not be parsed from JWT")
	}

	uID := userID.(float64)
	rID := roleID.(float64)

	jwtToken, err := s.repo.GetUserToken(ctx, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("GetUserToken err: %v", err)
	}

	type MyCustomClaims struct {
		UserID int `json:"user_id"`
		RoleID int `json:"role_id"`
		jwt.RegisteredClaims
	}

	newClaims := MyCustomClaims{
		int(uID),
		int(rID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	secretKey := []byte(s.jwtSecretKey)
	claimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	tokenString, err := claimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %v", err)
	}

	newRClaims := MyCustomClaims{
		int(uID),
		int(rID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(40 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	rClaimToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newRClaims)

	refreshTokenString, err := rClaimToken.SignedString(secretKey)
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %v", err)
	}

	jwtToken.Token = tokenString
	jwtToken.RefreshToken = refreshTokenString

	err = s.repo.UpdateUserToken(ctx, *jwtToken)
	if err != nil {
		return nil, fmt.Errorf("UpdateUserToken err: %v", err)
	}

	newToken := &JWTUserToken{
		Token:        jwtToken.Token,
		RefreshToken: jwtToken.RefreshToken,
	}

	return newToken, nil
}

func (s *Service) Register(ctx context.Context, request CreateUserRequest) (int, error) {
	request.Password = s.generatePassword(request.Password)

	user := transport.CreateUserRequest{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Login:     request.Login,
		Password:  request.Password,
	}

	resp, err := s.userGrpcTransport.CreateUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("CreateUser request err: %v", err)
	}

	randNum1 := rand.Intn(10)
	randNum2 := rand.Intn(10)
	randNum3 := rand.Intn(10)
	randNum4 := rand.Intn(10)

	msg := dto.UserCode{
		Email: request.Email,
		Code:  fmt.Sprintf("%d%d%d%d", randNum1, randNum2, randNum3, randNum4),
	}

	b, err := json.Marshal(&msg)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal UserCode err: %v", err)
	}

	s.userVerificationProducer.ProduceMessage(b)

	return int(resp.Id), nil
}

func (s *Service) ConfirmUser(ctx context.Context, request ConfirmUserRequest) error {
	res, err := s.redisCli.Get(ctx, request.Email).Result()
	if err != nil {
		return fmt.Errorf("failed to redis get err: %v", err)
	}

	if res != request.Code {
		return fmt.Errorf("wrong code")
	}

	err = s.userGrpcTransport.ConfirmUser(ctx, request.Email)
	if err != nil {
		return fmt.Errorf("failed to ConfirmUser err: %v", err)
	}

	return nil
}

func (s *Service) generatePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(s.passwordSecretKey))
	_, _ = hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *Service) validateToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method err: %v", t.Header["alg"])
		}

		return []byte(s.passwordSecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("parse token err: %v", err)
	}

	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
