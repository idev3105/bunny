package kafka

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/IBM/sarama/mocks"
)

func TestConsumer(t *testing.T) {
	t.Run("Consume messages successfully", func(t *testing.T) {
		// Create a mock Sarama consumer
		mockConsumer := mocks.NewConsumer(t, nil)

		// Set up topic metadata
		mockConsumer.SetTopicMetadata(map[string][]int32{
			"test-topic": {0},
		})

		mockPartitionConsumer := mockConsumer.ExpectConsumePartition("test-topic", 0, sarama.OffsetNewest)
		mockPartitionConsumer.ExpectMessagesDrainedOnClose()

		mockMessage := &sarama.ConsumerMessage{
			Value: []byte("test-value"),
			Topic: "test-topic",
		}
		mockPartitionConsumer.YieldMessage(mockMessage)

		// Create our Consumer with the NewConsumer method
		consumer, err := NewConsumer("localhost", 9092, "test-topic")
		if err != nil {
			t.Fatalf("Failed to create consumer: %v", err)
		}
		// Replace the client with our mock
		consumer.client = mockConsumer

		// Create a channel to receive messages
		messageChan := make(chan []byte, 1)

		// Start consuming in a goroutine
		ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Millisecond)
		defer cancel()

		go func() {
			err := consumer.Consume(ctx, func(message []byte) error {
				messageChan <- message
				return nil
			})
			if err != nil {
				t.Errorf("Consume returned an error: %v", err)
			}
		}()

		// Wait for a message
		select {
		case receivedMsg := <-messageChan:
			if !reflect.DeepEqual([]byte("test-value"), receivedMsg) {
				t.Errorf("Expected value %v, got %v", []byte("test-value"), receivedMsg)
			}
		case <-ctx.Done():
			t.Fatal("Timed out waiting for message")
		}

		if err := mockConsumer.Close(); err != nil {
			t.Errorf("Error closing the mock consumer: %v", err)
		}
	})

	t.Run("Handle context cancellation", func(t *testing.T) {
		// Create a mock Sarama consumer
		mockConsumer := mocks.NewConsumer(t, nil)

		mockConsumer.SetTopicMetadata(map[string][]int32{
			"test-topic-2": {0},
		})

		mockPartitionConsumer := mockConsumer.ExpectConsumePartition("test-topic-2", 0, sarama.OffsetNewest)
		mockPartitionConsumer.ExpectMessagesDrainedOnClose()

		mockMessage := &sarama.ConsumerMessage{
			Value: []byte("test-value"),
			Topic: "test-topic-2",
		}
		mockPartitionConsumer.YieldMessage(mockMessage)

		// Create our Consumer with the NewConsumer method
		consumer, err := NewConsumer("localhost", 9092, "test-topic-2")
		if err != nil {
			t.Fatalf("Failed to create consumer: %v", err)
		}
		// Replace the client with our mock
		consumer.client = mockConsumer

		ctx, cancel := context.WithCancel(context.Background())

		// Use a WaitGroup to ensure the Consume function has started
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			wg.Done() // Signal that the goroutine has started
			err := consumer.Consume(ctx, func(message []byte) error {
				// Process the message
				if !reflect.DeepEqual([]byte("test-value"), message) {
					t.Errorf("Expected value %v, got %v", []byte("test-value"), message)
				}
				// Cancel the context after processing the message
				cancel()
				return nil
			})
			if err != nil {
				t.Errorf("Consume returned an error: %v", err)
			}
		}()

		// Wait for the Consume function to start
		wg.Wait()

		// Wait a short time to allow message processing
		time.Sleep(100 * time.Millisecond)

		// Cancel the context if it hasn't been cancelled already
		cancel()

		// Wait for the mock consumer to be closed
		time.Sleep(100 * time.Millisecond)

		if err := mockConsumer.Close(); err != nil {
			t.Errorf("Error closing the mock consumer: %v", err)
		}
	})

	// Add more test cases as needed
}
