package auth

import (
	"context"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func (s *Service) GenerateToken(ctx context.Context, request GenerateTokenRequest) (*JWTUserToken, error) {
	user, err := s.userGrpcTransport.GetUserByLogin(ctx, request.Login)
	if err != nil {
		return nil, fmt.Errorf("failed to GetUserByLogin err: %v", err)
	}

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
