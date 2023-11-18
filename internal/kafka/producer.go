package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
)

type Producer struct {
	asyncProducer sarama.AsyncProducer
	topic         string
}

func NewProducer(cfg config.Kafka) (*Producer, error) {
	saramaCfg := sarama.NewConfig()

	asyncProducer, err := sarama.NewAsyncProducer(cfg.Brokers, saramaCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to NewAsyncProducer err: %v", err)
	}

	return &Producer{
		asyncProducer: asyncProducer,
		topic:         cfg.Producer.Topic,
	}, nil
}

func (p *Producer) ProduceMessage(message []byte) {
	msg := &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(message),
	}

	p.asyncProducer.Input() <- msg
}
