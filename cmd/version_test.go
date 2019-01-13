package cmd

import (
	"fmt"
	"testing"
)

func TestNewShowVeesion(t *testing.T) {
	versionCmd := NewShowVersion()
	stdout, stderr, err := executeCommand(versionCmd, "version")
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}

	expected := fmt.Sprintf("slackctl version: %s\n", Version)
	if stdout != expected {
		t.Errorf("subcommand version output is invalid:\ngot:\n%s\nexpected:\n%s", stdout, expected)
	}

	if len(stderr) != 0 {
		t.Errorf("expect stderr is nothing but got:\n%s", stderr)
	}
}
