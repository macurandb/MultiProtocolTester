package kafka

import (
	"github.com/macurandb/MultiProtocolTester/internal/protocols/kafka"
	"github.com/macurandb/MultiProtocolTester/services"
	"github.com/spf13/cobra"
	"log"
)

// kafkaReceiveCmd is used to receive messages from a Kafka topic.
var kafkaReceiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive messages from a Kafka topic",
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the config flag from the command
		configPath, _ := cmd.Flags().GetString("config")

		// If no config is provided, default to "config/kafka_config.yaml"
		if configPath == "" {
			configPath = "config/kafka_config.yaml"
		}

		// Load Kafka configuration
		config, err := kafka.LoadKafkaConfig(configPath)
		if err != nil {
			log.Fatalf("Failed to load Kafka configuration from %s: %v", configPath, err)
		}

		// Use the KafkaService to receive messages
		service := services.NewKafkaService()
		err = service.ReceiveMessages(config)
		if err != nil {
			log.Fatalf("Error receiving messages: %v", err)
		}
	},
}

func init() {
	kafkaReceiveCmd.Flags().StringP("config", "c", "", "Optional configuration file path")

	KafkaCmd.AddCommand(kafkaReceiveCmd)
}
