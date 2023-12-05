package repository

import (
	"context"
	"github.com/agadilkhan/pickup-point-service/internal/auth/entity"
	"gorm.io/gorm"
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

func (r *Repo) UpdateMessage(ctx context.Context, message entity.OutboxMessage) error {
	res := r.main.DB.WithContext(ctx).Where("code=?", message.Code).Updates(&message)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Repo) ProcessMessage(ctx context.Context, message entity.OutboxMessage) error {
	err := r.main.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Updates(message)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		res = tx.Where("id = ?", message.ID).Delete(message)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
