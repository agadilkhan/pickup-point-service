package repository

import (
	"context"

	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
)

func (r *Repo) SaveOutboxMessage(ctx context.Context, message entity.OutboxMessage) (int, error) {
	res := r.main.DB.WithContext(ctx).Create(&message)
	if res.Error != nil {
		return 0, res.Error
	}

	return message.ID, nil
}

func (r *Repo) GetUnProcessedMessages(ctx context.Context) (*[]entity.OutboxMessage, error) {
	var messages []entity.OutboxMessage

	res := r.replica.WithContext(ctx).Where("is_processed = false").Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}

	return &messages, nil
}

func (r *Repo) GetProcessedMessages(ctx context.Context) (*[]entity.OutboxMessage, error) {
	var messages []entity.OutboxMessage

	res := r.replica.WithContext(ctx).Where("is_processed = true").Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}

	return &messages, nil
}

func (r *Repo) UpdateMessage(ctx context.Context, code string) error {
	res := r.main.DB.Model(entity.OutboxMessage{}).WithContext(ctx).Where("code=?", code).Update("is_processed", "true")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Repo) DeleteMessage(ctx context.Context, message entity.OutboxMessage) error {
	res := r.main.DB.WithContext(ctx).Where("id = ?", message.ID).Delete(&message)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
