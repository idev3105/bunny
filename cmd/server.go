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
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := server.Create(cmd.Context())
			if err != nil {
				return err
			}
			err = s.Start()
			if err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
