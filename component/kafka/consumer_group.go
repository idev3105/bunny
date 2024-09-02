package kafka

import (
	"context"
	"errors"
	"fmt"

	"github.com/IBM/sarama"
	"org.idev.bunny/backend/common/logger"
)

type ConsumerGroup struct {
	Host          string
	Port          int32
	GroupID       string
	Topics        []string
	OnReceiveFunc func(ctx context.Context, value []byte) error
	Ready         chan bool

	consumerGroup sarama.ConsumerGroup
}

func NewConsumerGroup(host string, port int32, groupID string, topics []string, onReceiveFunc func(context.Context, []byte) error) (*ConsumerGroup, error) {
	addrs := []string{fmt.Sprintf("%s:%d", host, port)}

	cfg := sarama.NewConfig()
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	cg, err := sarama.NewConsumerGroup(addrs, groupID, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer group: %w", err)
	}

	return &ConsumerGroup{
		Host:          host,
		Port:          port,
		GroupID:       groupID,
		Topics:        topics,
		OnReceiveFunc: onReceiveFunc,
		Ready:         make(chan bool),
		consumerGroup: cg,
	}, nil
}

func (k *ConsumerGroup) Start(ctx context.Context) error {
	log := logger.New("ConsumerGroup", "Start")

	log.Info("Starting consumer in group: ", k.GroupID)
	for {
		if err := k.consumerGroup.Consume(ctx, k.Topics, k); err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}
			return fmt.Errorf("error from consumer: %w", err)
		}

		if ctx.Err() != nil {
			return nil
		}
		k.Ready = make(chan bool)
	}
}

func (k *ConsumerGroup) Close() error {
	log := logger.New("ConsumerGroup", "Close")
	log.Info("Closing consumer")
	if err := k.consumerGroup.Close(); err != nil {
		return fmt.Errorf("error closing consumer group: %w", err)
	}
	return nil
}

func (k *ConsumerGroup) Setup(sarama.ConsumerGroupSession) error {
	close(k.Ready)
	return nil
}

func (k *ConsumerGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (k *ConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log := logger.New("ConsumerGroup", "ConsumeClaim")
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Info("Message channel was closed")
				return nil
			}
			log.Infof("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)

			if err := k.OnReceiveFunc(session.Context(), message.Value); err != nil {
				return fmt.Errorf("error processing message: %w", err)
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
