package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdAuth() *cobra.Command {
	c := &cobra.Command{
		Use:   "auth",
		Short: "auth command",
		Long:  `This method checks authentication and tells "you" who you are, even if you might be a bot.`,
	}

	c.AddCommand(SubCmdAuthTest())
	return c
}

func SubCmdAuthTest() *cobra.Command {
	c := &cobra.Command{
		Use:   "test",
		Short: "Checks authentication & identity.",
		Long: `This method checks authentication and tells "you" who you are, even if you might be a bot.

You can also use this method to test whether Slack API authentication is functional.

API Reference is https://api.slack.com/methods/auth.test`,
		RunE: ShowAuthTest,
	}

	return c
}

func ShowAuthTest(c *cobra.Command, args []string) error {
	_, err := api.AuthTest()
	if err != nil {
		return fmt.Errorf("Calling auth.test is failed: %v", err)
	}
	c.Println("ok")
	return nil
}
