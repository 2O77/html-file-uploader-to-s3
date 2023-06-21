package kafkastorage

import (
	"fmt"
	"os"

	messagequeue "github.com/2o77/wope_case/API/internal/domain/message-queue"
	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
)

var kafkaTopicName string

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	kafkaTopicName = os.Getenv("KAFKA_TOPIC_NAME")
}

type KafkaStorage struct {
	producer sarama.SyncProducer
}

func NewKafka() (*KafkaStorage, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:29092"}, config)
	if err != nil {
		return nil, err
	}

	return &KafkaStorage{
		producer: producer,
	}, nil
}

func (kafkaStorage *KafkaStorage) PublishMessage(message messagequeue.Message) error {

	msg := &sarama.ProducerMessage{
		Topic: kafkaTopicName,
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
