package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
	"github.com/agadilkhan/pickup-point-service/internal/auth/repository"
	"github.com/agadilkhan/pickup-point-service/internal/auth/transport"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service struct {
	repo              repository.Repository
	jwtSecretKey      string
	passwordSecretKey string
	userTransport     *transport.UserTransport
}

func NewAuthService(
	repo repository.Repository,
	authConfig config.Auth,
	userTransport *transport.UserTransport,
) UseCase {
	return &Service{
		repo:              repo,
		jwtSecretKey:      authConfig.JWTSecretKey,
		passwordSecretKey: authConfig.PasswordSecretKey,
		userTransport:     userTransport,
	}
}

func (s *Service) GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JWTUserToken, error) {
	user, err := s.userTransport.GetUser(ctx, request.Login)
	if err != nil {
		return nil, fmt.Errorf("GetUser request err: %w", err)
	}

	generatedPassword := s.generatePassword(request.Password)
	if user.Password != generatedPassword {
		return nil, fmt.Errorf("password is wrong")
	}

	type MyCustomClaims struct {
		UserID int `json:"user_id"`
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		user.ID,
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
		user.ID,
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
		UserID:       user.ID,
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
	claims, err := s.ValidateToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("ValidateToken err: %v", err)
	}

	userID, ok := claims["user_id"]
	if !ok {
		return nil, fmt.Errorf("user_id could not be parsed from JWT")
	}

	id := userID.(int)

	type MyCustomClaims struct {
		UserID int `json:"user_id"`
		jwt.RegisteredClaims
	}

	newClaims := MyCustomClaims{
		id,
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
		id,
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

	userToken := entity.UserToken{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
		UserID:       id,
	}

	err = s.repo.UpdateUserToken(ctx, userToken)
	if err != nil {
		return nil, fmt.Errorf("UpdateUserToken err: %v", err)
	}

	jwtToken := &JWTUserToken{
		Token:        tokenString,
		RefreshToken: refreshTokenString,
	}

	return jwtToken, nil
}

func (s *Service) Register(ctx context.Context) (int, error) {
	return 0, nil
}

func (s *Service) generatePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(s.passwordSecretKey))
	_, _ = hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (s *Service) ValidateToken(tokenString string) (jwt.MapClaims, error) {
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
