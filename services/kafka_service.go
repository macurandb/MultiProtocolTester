package services

import (
	confluentKafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/macurandb/MultiProtocolTester/internal/protocols/kafka"
	"log"
)

// KafkaService handles the business logic for Kafka operations.
type KafkaService struct{}

// NewKafkaService creates a new KafkaService instance.
func NewKafkaService() *KafkaService {
	return &KafkaService{}
}

func (s *KafkaService) SendMessage(config *kafka.KafkaConfig, message string) error {
	// Create Kafka producer using Confluent Kafka Go package
	producer, err := confluentKafka.NewProducer(&confluentKafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers, // Ensure this is correctly passed
	})
	if err != nil {
		log.Printf("Failed to create producer: %v", err)
		return err
	}
	defer producer.Close()

	// Use internal Kafka sender to send the message
	sender := kafka.NewKafkaSender(producer, config.Topics.SendTopic)

	// Send the message using the sender
	err = sender.Send([]byte(message))
	if err != nil {
		log.Printf("Error while sending message: %v", err)
		return err
	}

	return nil
}

// ReceiveMessages handles receiving messages from Kafka.
func (s *KafkaService) ReceiveMessages(config *kafka.KafkaConfig) error {
	// Delegate to the internal Kafka logic for receiving messages
	kafka.ReceiveMessagesWithConfig(config)
	return nil
}
