package consumer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/IBM/sarama"
	"github.com/agadilkhan/pickup-point-service/internal/auth/controller/consumer/dto"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type UserVerificationCallback struct {
	logger   *zap.SugaredLogger
	redisCli *redis.Client
}

func NewUserVerificationCallback(logger *zap.SugaredLogger, redisCli *redis.Client) *UserVerificationCallback {
	return &UserVerificationCallback{
		logger:   logger,
		redisCli: redisCli,
	}
}

func (c *UserVerificationCallback) Callback(message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError) {
	for {
		select {
		case msg := <-message:
			var userCode dto.UserCode

			err := json.Unmarshal(msg.Value, &userCode)
			if err != nil {
				c.logger.Errorf("failed to unmarshal record value err: %v", err)
			} else {
				c.logger.Infof("user code: %v", userCode)

				err = c.redisCli.Set(context.Background(), userCode.Email, userCode.Code, 10*time.Minute).Err()
				if err != nil {
					c.logger.Errorf("failed to save record value in cache err: %v", err)
				}
			}
		case err := <-error:
			c.logger.Errorf("failed consume err: %v", err)
		}
	}
}
