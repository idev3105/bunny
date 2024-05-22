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
	OnReceiveFunc func([]byte) error
	Ready         chan bool

	consumerGroup sarama.ConsumerGroup
}

func NewConsumerGroup(host string, port int32, groupID string, topics []string, onReceiveFunc func([]byte) error) (*ConsumerGroup, error) {
	addrs := []string{fmt.Sprintf("%v:%v", host, port)}

	cfg := sarama.NewConfig()
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	cg, err := sarama.NewConsumerGroup(addrs, groupID, cfg)
	if err != nil {
		return nil, err
	}
	return &ConsumerGroup{
		Host:          host,
		Port:          port,
		GroupID:       groupID,
		Topics:        topics,
		OnReceiveFunc: onReceiveFunc,
		Ready:         make(chan bool),
		consumerGroup: cg,
	}, err
}

func (k *ConsumerGroup) Start(ctx context.Context) error {
	log := logger.New("ConsumerGroup", "Start")

	log.Info("Start this consumer in group: ", k.GroupID)
	for {
		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		if err := k.consumerGroup.Consume(ctx, k.Topics, k); err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}
			log.Errorf("Error from consumer: %v", err)
			return err
		}
		// check if context was cancelled, signaling that the consumer should stop
		if ctx.Err() != nil {
			return nil
		}
		k.Ready = make(chan bool)
	}
}

func (k *ConsumerGroup) Close() error {
	log := logger.New("ConsumerGroup", "Close")
	log.Info("Close this consumer")
	if err := k.consumerGroup.Close(); err != nil {
		log.Errorf("Error closing consumer group: %v", err)
		return err
	}
	return nil
}

// Implement sarama.ConsumerGroupHandler interface
func (k *ConsumerGroup) Setup(session sarama.ConsumerGroupSession) error {
	close(k.Ready)
	return nil
}

// Implement sarama.ConsumerGroupHandler interface
func (k *ConsumerGroup) Cleanup(session sarama.ConsumerGroupSession) error {
	// TODO: Implement here
	return nil
}

// Implement sarama.ConsumerGroupHandler interface
func (k *ConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	logger := logger.New("ConsumerGroup", "ConsumeClaim")
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				logger.Info("Message channel was closed")
				return nil
			}
			logger.Infof("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}
