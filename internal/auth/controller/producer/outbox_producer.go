package producer

import (
	"context"
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
	workers  int
}

func NewOutboxProducer(
	producer *kafka.Producer,
	repo repository.Repository,
	interval time.Duration,
	logger *zap.SugaredLogger,
	workers int,
) *OutboxProducer {
	return &OutboxProducer{
		producer: producer,
		repo:     repo,
		interval: interval,
		logger:   logger,
		workers:  workers,
	}
}

func (op *OutboxProducer) Run(ctx context.Context) {
	ticker := time.NewTimer(op.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			op.task1(ctx)

			op.task2(ctx)

			ticker.Reset(op.interval)

		case <-ctx.Done():
			return
		}
	}
}
