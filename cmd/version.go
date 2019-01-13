package cmd

import "github.com/spf13/cobra"

var Version = "0.0.1"

func NewShowVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("slackctl version: %s\n", Version)
		},
	}
	return cmd
}
