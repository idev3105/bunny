package cmd

import (
	"github.com/spf13/cobra"
	exampleconsumer "org.idev.bunny/backend/consumer/example"
)

func NewConsumerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "consumer",
		Short: "consumer command",
		Long:  "consumer command",
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 1 {
				panic("consumer name is required")
			}

			consumerName := args[0]
			if consumerName == "example" {
				c, err := exampleconsumer.New("test-group", []string{"test1", "test2"})
				if err != nil {
					panic(err)
				}
				if err := c.Start(); err != nil {
					panic(err)
				}
			}
		},
	}
}
