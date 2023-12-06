package producer

import (
	"context"
	"encoding/json"

	"github.com/agadilkhan/pickup-point-service/internal/auth/controller/consumer/dto"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
)

func (op *OutboxProducer) task1(ctx context.Context) {
	messages, err := op.repo.GetUnProcessedMessages(ctx)
	if err != nil {
		op.logger.Errorf("failed to GetUnprocessedMessages")
	}

	jobs := make(chan entity.OutboxMessage, len(*messages))
	results := make(chan entity.OutboxMessage, len(*messages))

	for i := 0; i < op.workers; i++ {
		go op.processMessage(ctx, jobs, results)
	}

	for _, msg := range *messages {
		jobs <- msg
	}

	for i := 0; i < len(*messages); i++ {
		msg := <-results
		kafkaMessage := dto.UserCode{
			Code:  msg.Code,
			Email: msg.UserEmail,
		}

		b, err := json.Marshal(&kafkaMessage)
		if err != nil {
			op.logger.Errorf("failed to Marshal err: %v", err)
			continue
		}

		op.producer.ProduceMessage(b)
	}

	close(results)
	close(jobs)
}

func (op *OutboxProducer) processMessage(
	ctx context.Context,
	jobs <-chan entity.OutboxMessage,
	results chan<- entity.OutboxMessage,
) {

	for msg := range jobs {
		msg.IsProcessed = true
		err := op.repo.UpdateMessage(ctx, msg.Code)
		if err != nil {
			op.logger.Errorf("failed to UpdateMessage err: %v", err)
			continue
		}
		op.logger.Info("message updated")
		results <- msg
	}
}
