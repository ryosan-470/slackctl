package cmd

import (
	"bytes"

	"github.com/spf13/cobra"
)

func executeCommand(c *cobra.Command, args ...string) (stdout, stderr string, err error) {
	out, outerr := new(bytes.Buffer), new(bytes.Buffer)
	c.SetArgs(args)
	if err := Execute(c, out, outerr); err != nil {
		return "", "", err
	}
	return out.String(), outerr.String(), nil
}
