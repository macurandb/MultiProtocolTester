package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

// kafkaSender struct is used to send messages to Kafka.
type kafkaSender struct {
	producer *kafka.Producer
	topic    string
}

// NewKafkaSender creates a new Kafka sender.
func NewKafkaSender(producer *kafka.Producer, topic string) *kafkaSender {
	return &kafkaSender{
		producer: producer,
		topic:    topic,
	}
}

// Send sends a message to the Kafka topic.
func (k *kafkaSender) Send(message []byte) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &k.topic,
			Partition: kafka.PartitionAny,
		},
		Value: message,
	}

	err := k.producer.Produce(msg, nil)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	// Wait for delivery report from producer.Events()
	e := <-k.producer.Events()
	m, ok := e.(*kafka.Message)
	if !ok || m.TopicPartition.Error != nil {
		if m.TopicPartition.Error != nil {
			log.Printf("Message delivery failed: %v", m.TopicPartition.Error)
		}
		return m.TopicPartition.Error
	}

	log.Printf("Message delivered to partition %d at offset %v", m.TopicPartition.Partition, m.TopicPartition.Offset)
	return nil
}
