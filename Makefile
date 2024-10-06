# Go settings
BINARY_NAME=prototester
DOCKER_COMPOSE_FILE=docker-compose.yml
KAFKA_CONFIG=config/kafka_config.yaml

# Default target
.PHONY: all
all: build

# Build the Go binary
.PHONY: build
build:
	@echo "Building the CLI application..."
	go build -o $(BINARY_NAME) main.go

# Run tests (add your Go tests if you have any)
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# Start the protocol services (e.g., Kafka, RabbitMQ) using Docker Compose
.PHONY: start-services
start-services:
	@echo "Starting protocol services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop the protocol services using Docker Compose
.PHONY: stop-services
stop-services:
	@echo "Stopping protocol services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Generic send command for any protocol
.PHONY: %/send
%/send:
	@echo "Sending a message using $*..."
	./$(BINARY_NAME) $* send --message "Hello from $*!" --config $(KAFKA_CONFIG)

# Generic receive command for any protocol
.PHONY: %/receive
%/receive:
	@echo "Receiving messages using $*..."
	./$(BINARY_NAME) $* receive --config $(KAFKA_CONFIG)

# Clean up build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Rebuild the binary
.PHONY: rebuild
rebuild: clean build
