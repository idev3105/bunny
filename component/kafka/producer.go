package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	Host     string
	Port     int32
	producer sarama.SyncProducer
}

func NewProducer(host string, port int32) (*Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	p, err := sarama.NewSyncProducer([]string{fmt.Sprintf("%v:%v", host, port)}, cfg)
	if err != nil {
		return nil, err
	}
	return &Producer{
		Host:     host,
		Port:     port,
		producer: p,
	}, nil
}

func (k *Producer) Send(messages []Message) error {
	msg := make([]*sarama.ProducerMessage, len(messages))
	for i, m := range messages {
		msg[i] = &sarama.ProducerMessage{
			Key:   sarama.ByteEncoder(m.Key),
			Topic: m.Topic,
			Value: sarama.ByteEncoder(m.Value),
		}
	}
	err := k.producer.SendMessages(msg)
	return err
}

func (k *Producer) SendAsync(messages []Message) error {
	panic("not implemented")
}

func (k *Producer) Close() error {
	return k.producer.Close()
}
