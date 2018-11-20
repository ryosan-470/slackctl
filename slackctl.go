package slackctl

import (
	"github.com/nlopes/slack"
	"github.com/urfave/cli"
)

type SlackCtl struct {
	Client  *slack.Client
	Context *cli.Context
}
