package kafkastorage

import (
	messagequeue "github.com/2o77/wope_case/API/internal/domain/message-queue"
	"github.com/Shopify/sarama"
)

type KafkaStorage struct {
	producer sarama.SyncProducer
	consumer sarama.Consumer
}

func NewKafka() (*KafkaStorage, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:29092"}, config)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumer([]string{"localhost:29092"}, config)
	if err != nil {
		return nil, err
	}

	return &KafkaStorage{
		producer: producer,
		consumer: consumer,
	}, nil
}

func (kafkaStorage *KafkaStorage) PublishMessage(message messagequeue.Message) error {

	msg := &sarama.ProducerMessage{
		Topic: "wope-case-topic",
		Value: sarama.StringEncoder(message.FilePath),
	}

	_, _, err := kafkaStorage.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (kafkaStorage *KafkaStorage) Close() error {
	err := kafkaStorage.producer.Close()
	if err != nil {
		return err
	}

	return nil
}
