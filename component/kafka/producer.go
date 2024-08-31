package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

// Producer represents a Kafka producer
type Producer struct {
	syncProducer  sarama.SyncProducer
	asyncProducer sarama.AsyncProducer
}

// NewProducer creates a new Kafka producer
func NewProducer(host string, port int32) (*Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	address := fmt.Sprintf("%s:%d", host, port)

	syncProducer, err := sarama.NewSyncProducer([]string{address}, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka sync producer: %w", err)
	}

	asyncProducer, err := sarama.NewAsyncProducer([]string{address}, cfg)
	if err != nil {
		syncProducer.Close() // Close the sync producer if async producer creation fails
		return nil, fmt.Errorf("failed to create Kafka async producer: %w", err)
	}

	return &Producer{syncProducer: syncProducer, asyncProducer: asyncProducer}, nil
}

// SendSync sends multiple messages to Kafka synchronously
func (p *Producer) SendSync(messages []Message) error {
	producerMessages := make([]*sarama.ProducerMessage, len(messages))
	for i, m := range messages {
		producerMessages[i] = &sarama.ProducerMessage{
			Key:   sarama.ByteEncoder(m.Key),
			Topic: m.Topic,
			Value: sarama.ByteEncoder(m.Value),
		}
	}
	return p.syncProducer.SendMessages(producerMessages)
}

// SendAsync sends multiple messages to Kafka asynchronously
func (p *Producer) SendAsync(messages []Message) error {
	for _, m := range messages {
		producerMessage := &sarama.ProducerMessage{
			Key:   sarama.ByteEncoder(m.Key),
			Topic: m.Topic,
			Value: sarama.ByteEncoder(m.Value),
		}
		p.asyncProducer.Input() <- producerMessage
	}
	return nil
}

// Close closes both the sync and async Kafka producers
func (p *Producer) Close() error {
	var errs []error

	if err := p.syncProducer.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close sync producer: %w", err))
	}

	if err := p.asyncProducer.Close(); err != nil {
		errs = append(errs, fmt.Errorf("failed to close async producer: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing producers: %v", errs)
	}

	return nil
}
