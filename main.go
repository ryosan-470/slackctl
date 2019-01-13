package main

import (
	"os"

	"github.com/ryosan-470/slackctl/cmd"
)

func main() {
	rootCmd := cmd.NewCmdRoot()
	if err := cmd.Execute(rootCmd, os.Stdout, os.Stderr); err != nil {
		os.Exit(1)
	}
}
