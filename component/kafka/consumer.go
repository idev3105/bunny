package kafka

import (
	"context"
	"fmt"
	"sync"

	"github.com/IBM/sarama"
	"org.idev.bunny/backend/common/logger"
)

type Consumer struct {
	Host   string
	Port   int32
	Topic  string
	client sarama.Consumer
}

func NewConsumer(host string, port int32, topic string) (*Consumer, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	client, err := sarama.NewConsumer([]string{addr}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	return &Consumer{
		Host:   host,
		Port:   port,
		Topic:  topic,
		client: client,
	}, nil
}

func (c *Consumer) Consume(ctx context.Context, onReceive func(message []byte) error) error {
	log := logger.New("Consumer", "Consume")

	partitions, err := c.client.Partitions(c.Topic)
	if err != nil {
		return fmt.Errorf("failed to get partitions: %w", err)
	}

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	for _, partition := range partitions {
		pc, err := c.client.ConsumePartition(c.Topic, partition, sarama.OffsetNewest)
		if err != nil {
			return fmt.Errorf("failed to start consumer for partition %d: %w", partition, err)
		}

		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			defer pc.Close()

			for {
				select {
				case msg, ok := <-pc.Messages():
					if !ok {
						log.Info("Messages channel closed")
						return
					}
					if msg == nil {
						log.Debug("Received nil message")
						continue
					}
					log.Debug("Received message", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset, "value", string(msg.Value))
					if err := onReceive(msg.Value); err != nil {
						log.Error("Error processing message", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset, "error", err)
					}
				case err := <-pc.Errors():
					log.Error("Error while consuming: ", err)
				case <-ctx.Done():
					log.Info("Context cancelled, stopping consumer for partition", partition)
					return
				}
			}
		}(pc)
	}

	log.Info("Started consuming from topic: ", c.Topic)

	// Wait for context cancellation
	<-ctx.Done()
	log.Info("Context cancelled, stopping consumer")

	// Wait for all goroutines to finish
	wg.Wait()
	log.Info("All partition consumers closed")

	return nil
}

func (c *Consumer) Close() error {
	return c.client.Close()
}
