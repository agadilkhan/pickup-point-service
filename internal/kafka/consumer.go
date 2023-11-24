package kafka

import (
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"github.com/agadilkhan/pickup-point-service/internal/auth/config"
	"go.uber.org/zap"
)

type ConsumerCallback interface {
	Callback(message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError)
}

type Consumer struct {
	logger   *zap.SugaredLogger
	topics   []string
	master   sarama.Consumer
	callback ConsumerCallback
}

func NewConsumer(
	logger *zap.SugaredLogger,
	cfg config.Kafka,
	callback ConsumerCallback,
) (*Consumer, error) {
	saramaCfg := sarama.NewConfig()
	saramaCfg.ClientID = "go-kafka-consumer"
	saramaCfg.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(cfg.Brokers, saramaCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create NewConsumer err: %v", err)
	}

	return &Consumer{
		logger:   logger,
		topics:   cfg.Consumer.Topics,
		master:   master,
		callback: callback,
	}, nil
}

func (c *Consumer) Start() {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)

	for _, topic := range c.topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}

		partitions, _ := c.master.Partitions(topic)

		consumer, err := c.master.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
		if err != nil {
			c.logger.Errorf("Topic %v Partitions: %v, err: %v", topic, partitions, err)
			continue
		}

		c.logger.Info("Start consuming topic", topic)

		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerErr := <-consumer.Errors():
					errors <- consumerErr
				case msg := <-consumer.Messages():
					consumers <- msg
				}
			}
		}(topic, consumer)
	}

	c.callback.Callback(consumers, errors)
}
