package admin

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/admin/entity"
	"github.com/agadilkhan/pickup-point-service/internal/admin/transport"
)

type Service struct {
	passwordSecretKey string
	userGrpcTransport *transport.UserGrpcTransport
}

func NewService(passwordSecretKey string, userGrpcTransport *transport.UserGrpcTransport) *Service {
	return &Service{
		passwordSecretKey: passwordSecretKey,
		userGrpcTransport: userGrpcTransport,
	}
}

func (s *Service) GetUsers(ctx context.Context) (*[]entity.User, error) {
	return nil, nil
}

func (s *Service) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	resp, err := s.userGrpcTransport.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to GetUserByID err: %v", err)
	}

	user := entity.User{
		ID:          int(resp.Id),
		RoleID:      int(resp.RoleId),
		FirstName:   resp.FirstName,
		LastName:    resp.LastName,
		Email:       resp.Email,
		Phone:       resp.Phone,
		Login:       resp.Login,
		Password:    resp.Password,
		IsConfirmed: resp.IsConfirmed,
	}

	return &user, nil
}

func (s *Service) UpdateUser(ctx context.Context, request UpdateUserRequest) (*entity.User, error) {
	request.Password = s.generatePassword(request.Password)

	updateRequest := transport.UpdateUserRequest{
		ID:        request.ID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Login:     request.Login,
		Password:  request.Password,
	}

	resp, err := s.userGrpcTransport.UpdateUser(ctx, updateRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to UpdateUser err: %v", err)
	}

	return &entity.User{
		ID:          int(resp.Id),
		RoleID:      int(resp.RoleId),
		FirstName:   resp.FirstName,
		LastName:    resp.LastName,
		Email:       resp.Email,
		Phone:       resp.Phone,
		Login:       resp.Login,
		Password:    resp.Password,
		IsConfirmed: resp.IsConfirmed,
	}, nil
}

func (s *Service) DeleteUser(ctx context.Context, id int) (int, error) {
	resp, err := s.userGrpcTransport.DeleteUser(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("failed to DeleteUser err: %v", err)
	}

	return int(resp.Id), nil
}

func (s *Service) generatePassword(password string) string {
	hash := hmac.New(sha256.New, []byte(s.passwordSecretKey))
	_, _ = hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
