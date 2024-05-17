package main

import (
	"github.com/spf13/cobra"
	"org.idev.bunny/backend/cmd"
	_ "org.idev.bunny/backend/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
func main() {
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
	rootCmd.Execute()
}
