package cmd

import (
	"io"
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

func Execute(cmd *cobra.Command, stdout, stderr io.Writer) error {
	cmd.SetOutput(stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(stderr)
		cmd.Println(err)
		return err
	}
	return nil
}
