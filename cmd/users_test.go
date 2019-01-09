package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandUsersRoot(t *testing.T) {
	buf := new(bytes.Buffer)
	c := NewCmdUsers()
	c.SetOutput(buf)
	c.SetArgs([]string{"users"})
	Execute()

	got := buf.String()
	expected := "hi"
	assert.Equal(t, expected, got)
}
