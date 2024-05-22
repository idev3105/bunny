package exampleconsumer

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"

	"org.idev.bunny/backend/common/logger"
)

func (ec *ExampleConsumer) Start(ctx context.Context) error {

	log := logger.New("ExampleConsumer", "Start consumer")
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := ec.consumerGroup.Start(ctx); err != nil {
			log.Fatalf("Shutting down consumer %v", err)
		}
	}()

	// wait for the consumer to be ready
	<-ec.consumerGroup.Ready
	log.Info("Consumer up and running!...")
	<-ctx.Done()

	// wait all job to finish before closing the consumer
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	wg.Wait()
	if err := ec.consumerGroup.Close(); err != nil {
		log.Errorf("Error closing consumer: %v", err)
		return err
	}

	return nil
}
