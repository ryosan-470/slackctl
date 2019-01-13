package cmd

import (
	"os"

	"github.com/nlopes/slack"
	"github.com/spf13/cobra"
)

var api *slack.Client

func init() {
	cobra.OnInitialize(initConfig)
}

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "slackctl",
		Short: "slackctl is a management tool for Slack using CLI",
		Long:  "slackctl is a management tool for Slack using CLI",
		Run:   runHelp,
	}
	cmd.AddCommand(NewCmdUsers())
	cmd.AddCommand(NewShowVersion())
	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	// snip
}

func initConfig() {
	token := os.Getenv("SLACK_TOKEN")
	api = slack.New(token)
}

func Execute() {
	cmd := NewCmdRoot()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
