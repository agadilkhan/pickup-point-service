package producer

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
)

func (op *OutboxProducer) task2(ctx context.Context) {
	messages, err := op.repo.GetProcessedMessages(ctx)
	if err != nil {
		op.logger.Errorf("failed to GetProcessedMessages err: %v", err)
	}

	jobs := make(chan entity.OutboxMessage, len(*messages))
	results := make(chan error, len(*messages))

	for i := 0; i < op.workers; i++ {
		go op.deleteMessage(ctx, jobs, results)
	}

	for _, msg := range *messages {
		jobs <- msg
	}

	for i := 0; i < len(*messages); i++ {
		err = <-results
		if err != nil {
			op.logger.Errorf("failed to DeleteMessage err: %v", err)
		} else {
			op.logger.Info("message deleted")
		}
	}

	close(results)
	close(jobs)
}

func (op *OutboxProducer) deleteMessage(
	ctx context.Context,
	jobs <-chan entity.OutboxMessage,
	results chan<- error,
) {

	for msg := range jobs {
		err := op.repo.DeleteMessage(ctx, msg)
		results <- err
	}

}
