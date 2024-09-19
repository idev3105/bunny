package main

import (
	"os"

	"github.com/spf13/cobra"
	"org.idev.bunny/backend/cmd"
	"org.idev.bunny/backend/common/errors"
	"org.idev.bunny/backend/common/logger"
	_ "org.idev.bunny/backend/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
func main() {
	log := logger.New("main", "Root")

	rootCmd := &cobra.Command{
		Use:   "",
		Short: "help",
		Long:  "show help menu",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.AddCommand(cmd.NewServerCommand())
	rootCmd.AddCommand(cmd.NewConsumerCmd())

	log.Infof("Application run on PID: %v", os.Getpid())

	if err := rootCmd.Execute(); err != nil {
		log.Error(errors.ToString(err))
		panic(err)
	}
}
