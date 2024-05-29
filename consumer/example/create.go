package exampleconsumer

import (
	"fmt"

	"org.idev.bunny/backend/app"
	"org.idev.bunny/backend/component/kafka"
)

func New(groupID string, topics []string) (*ExampleConsumer, error) {
	appConfig, err := app.LoadConfig()
	if err != nil {
		return nil, err
	}

	cg, err := kafka.NewConsumerGroup(appConfig.KafkaHost, appConfig.KafkaPort, groupID, topics, func(msg []byte) error {
		// TODO implement here
		fmt.Println(string(msg))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &ExampleConsumer{consumerGroup: cg}, nil
}
