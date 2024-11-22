package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

// ReceiveMessagesWithConfig receives messages from Kafka using the provided configuration.
func ReceiveMessagesWithConfig(config *KafkaConfig) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
		"group.id":          config.GroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Subscribe to the configured topic
	err = consumer.SubscribeTopics([]string{config.Topics.ReceiveTopic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic %s: %v", config.Topics.ReceiveTopic, err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message from topic %s: %s", *msg.TopicPartition.Topic, string(msg.Value))
		} else {
			log.Printf("Error receiving message: %v", err)
		}
	}
}
