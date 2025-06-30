package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)
type KafkaClient struct {
	Consumer *Consumer
	Producer *Producer
}
type Producer struct {
	producer sarama.SyncProducer
}

type Consumer struct {
	consumer sarama.Consumer
}
func NewClient(brokers []string) (*KafkaClient, error) {
	producer, err := NewProducer(brokers)
	if err != nil{
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	consumer, err := NewConsumer(brokers)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}
	return &KafkaClient{
		Producer: producer,
		Consumer: consumer,
	}, nil
}
func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	return &Producer{producer: producer}, nil
}

func NewConsumer(brokers []string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	return &Consumer{consumer: consumer}, nil
}