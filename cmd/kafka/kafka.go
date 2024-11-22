package kafka

import (
	"github.com/spf13/cobra"
)

// kafkaCmd is the parent command for Kafka-related operations.
var KafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Interact with Kafka for sending and receiving messages",
	Long:  "Use the Kafka subcommands to send and receive messages through Kafka topics.",
}
