package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/ryosan-470/slackctl"
	"github.com/urfave/cli"
)

const (
	name        = "slackctl"
	description = "Management tool for Slack"
	version     = "0.0.1"
)

var slackToken = os.Getenv("SLACK_TOKEN")

func main() {
	s := &slackctl.SlackCtl{
		Client: slack.New(slackToken),
	}

	app := cli.NewApp()
	app.Name = name
	app.Usage = description
	app.Version = version
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name:  "channel",
			Usage: "Get info on Slack channels, create or archive channels, invite users",
			Flags: []cli.Flag{},
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "create channel",
					Action: func(c *cli.Context) error {
						s.Context = c
						return s.CreateChannel()
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
