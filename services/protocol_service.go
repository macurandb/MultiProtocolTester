package services

import (
	"fmt"
	"log"
	"time"
)

// ProtocolService provides protocol-agnostic utility functions.
type ProtocolService struct{}

// Retry performs a retry operation for sending messages.
func (s *ProtocolService) Retry(operation func() error, retries int, delay time.Duration) error {
	for i := 0; i < retries; i++ {
		err := operation()
		if err == nil {
			return nil
		}
		log.Printf("Operation failed, retrying... (%d/%d)", i+1, retries)
		time.Sleep(delay)
	}
	return fmt.Errorf("operation failed after %d retries", retries)
}
