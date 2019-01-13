package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNewShowVeesion(t *testing.T) {
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	rootCmd := NewShowVersion()
	rootCmd.SetArgs([]string{"version"})
	err := Execute(rootCmd, stdout, stderr)
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}

	expected := fmt.Sprintf("slackctl version: %s\n", Version)
	if stdout.String() != expected {
		t.Errorf("subcommand version output is invalid:\ngot:\n%s\nexpected:\n%s", stdout.String(), expected)
	}

	if stderr.Len() != 0 {
		t.Errorf("expect stderr is nothing but got: %v", stderr)
	}
}
