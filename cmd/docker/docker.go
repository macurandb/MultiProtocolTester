package docker

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// dockerCmd is the parent command for managing Docker services.
var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Manage protocol services via Docker Compose",
	Long:  "This command allows you to start, stop, and manage the protocol services (Kafka, RabbitMQ, etc.) via Docker Compose.",
}

// startCmd starts the Docker Compose services.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start protocol services",
	Run: func(cmd *cobra.Command, args []string) {
		err := runDockerComposeCommand("up", "-d")
		if err != nil {
			log.Fatalf("Failed to start services: %v", err)
		}
		fmt.Println("Services started successfully.")
	},
}

// stopCmd stops the Docker Compose services.
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop protocol services",
	Run: func(cmd *cobra.Command, args []string) {
		err := runDockerComposeCommand("down")
		if err != nil {
			log.Fatalf("Failed to stop services: %v", err)
		}
		fmt.Println("Services stopped successfully.")
	},
}

// statusCmd checks the status of the Docker Compose services.
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check the status of protocol services",
	Run: func(cmd *cobra.Command, args []string) {
		err := runDockerComposeCommand("ps")
		if err != nil {
			log.Fatalf("Failed to check status: %v", err)
		}
	},
}

// runDockerComposeCommand runs Docker Compose commands from the CLI.
func runDockerComposeCommand(args ...string) error {
	cmd := exec.Command("docker-compose", args...)
	cmd.Stdout = cmd.Stderr // Show output on the terminal
	return cmd.Run()
}

func init() {
	DockerCmd.AddCommand(startCmd)
	DockerCmd.AddCommand(stopCmd)
	DockerCmd.AddCommand(statusCmd)
}
