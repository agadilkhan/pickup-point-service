package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
	"math/rand"

	"github.com/agadilkhan/pickup-point-service/internal/auth/controller/consumer/dto"
	"github.com/agadilkhan/pickup-point-service/internal/auth/transport"
)

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

	outboxMessage := entity.OutboxMessage{
		UserEmail: msg.Email,
		Code:      msg.Code,
	}

	_, err = s.repo.SaveOutboxMessage(ctx, outboxMessage)
	if err != nil {
		return 0, fmt.Errorf("failed to SaveOutboxMessage err: %v", err)
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
