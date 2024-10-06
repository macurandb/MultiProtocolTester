package cmd

import (
	"fmt"
	"github.com/macurandb/MultiProtocolTester/cmd/docker"
	"github.com/macurandb/MultiProtocolTester/cmd/kafka"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd is the base command when no subcommands are called.
var rootCmd = &cobra.Command{
	Use:   "prototester",
	Short: "CLI for interacting with multiple protocols",
	Long:  "This CLI allows you to send and receive messages through different protocols like Kafka, with support for additional protocols in the future.",
}

// Execute runs the root command and any subcommands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Initialize the CLI with protocol-specific commands.
func init() {
	rootCmd.AddCommand(kafka.KafkaCmd)
	rootCmd.AddCommand(docker.DockerCmd)
}
