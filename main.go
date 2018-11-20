package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	name        = "slackctl"
	description = "Management tool for Slack"
	version     = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Usage = description
	app.Version = version
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
