package cmd

import (
	"github.com/spf13/cobra"
	"org.idev.bunny/backend/api/server"
)

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "start server",
		Long:  "start server",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
	return cmd
}

func startServer() {
	s := server.Create()
	err := s.Start(":3000")
	if err != nil {
		panic(err)
	}
}
