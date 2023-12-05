package producer

import (
	"context"
	"encoding/json"
	"github.com/agadilkhan/pickup-point-service/internal/auth/controller/consumer/dto"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
	"github.com/agadilkhan/pickup-point-service/internal/auth/repository"
	"github.com/agadilkhan/pickup-point-service/internal/kafka"
	"go.uber.org/zap"
	"time"
)

type OutboxProducer struct {
	producer *kafka.Producer
	repo     repository.Repository
	interval time.Duration
	logger   *zap.SugaredLogger
}

func NewOutboxProducer(producer *kafka.Producer, repo repository.Repository, interval time.Duration, logger *zap.SugaredLogger) *OutboxProducer {
	return &OutboxProducer{
		producer: producer,
		repo:     repo,
		interval: interval,
		logger:   logger,
	}
}

func (op *OutboxProducer) Run(ctx context.Context) {
	ticker := time.NewTimer(op.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			messages, err := op.repo.GetUnProcessedMessages(ctx)
			if err != nil {
				op.logger.Errorf("failed to GetUnprocessedMessages")
			}

			results := make(chan error, len(*messages))

			for _, msg := range *messages {
				go op.ProcessMessage(msg, results)
			}

			for err := range results {
				if err != nil {
					op.logger.Errorf("failed to ProccessMessage")
				}
			}

		case <-ctx.Done():
			return
		}
	}
}

func (op *OutboxProducer) ProcessMessage(message entity.OutboxMessage, results chan<- error) {
	err := op.repo.ProcessMessage(context.Background(), message)
	results <- err
	if err == nil {
		msg := dto.UserCode{
			Code:  message.Code,
			Email: message.UserEmail,
		}

		b, err := json.Marshal(&msg)
		if err != nil {
			op.logger.Errorf("failed to marshal")
		}
		op.producer.ProduceMessage(b)
	}

}
