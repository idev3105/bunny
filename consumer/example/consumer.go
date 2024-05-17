package exampleconsumer

import "org.idev.bunny/backend/component/kafka"

type ExampleConsumer struct {
	consumerGroup *kafka.ConsumerGroup
}
