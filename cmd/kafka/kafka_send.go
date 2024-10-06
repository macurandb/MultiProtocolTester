package kafka

import (
	"github.com/macurandb/MultiProtocolTester/internal/protocols/kafka"
	"github.com/macurandb/MultiProtocolTester/services"
	"github.com/spf13/cobra"
	"log"
)

// kafkaSendCmd is used to send messages to a Kafka topic.
var kafkaSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a Kafka topic",
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the message and config flag from the command
		message, _ := cmd.Flags().GetString("message")
		configPath, _ := cmd.Flags().GetString("config")

		if configPath == "" {
			configPath = "config/kafka_config.yaml"
		}

		config, err := kafka.LoadKafkaConfig(configPath)
		if err != nil {
			log.Fatalf("Failed to load Kafka configuration from %s: %v", configPath, err)
		}

		service := services.NewKafkaService()
		err = service.SendMessage(config, message)
		if err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
	},
}

func init() {
	// Define the message flag (required)
	kafkaSendCmd.Flags().StringP("message", "m", "", "Message to send")
	kafkaSendCmd.MarkFlagRequired("message")

	// Define an optional config flag for specifying a custom config file
	kafkaSendCmd.Flags().StringP("config", "c", "", "Optional configuration file path")

	// Add the command to the Kafka root command
	KafkaCmd.AddCommand(kafkaSendCmd)
}
