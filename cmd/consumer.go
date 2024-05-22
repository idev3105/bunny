package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	exampleconsumer "org.idev.bunny/backend/consumer/example"
)

func NewConsumerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "consumer",
		Short: "consumer command",
		Long:  "consumer command",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				return errors.New("consumer name is required")
			}

			consumerName := args[0]
			if consumerName == "example" {
				c, err := exampleconsumer.New("test-group", []string{"test1", "test2"})
				if err != nil {
					return err
				}
				if err := c.Start(cmd.Context()); err != nil {
					return err
				}
			} else {
				return errors.New("consumer not found")
			}

			return nil
		},
	}
}
