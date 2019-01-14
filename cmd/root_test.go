package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/nlopes/slack"
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

func dummySlackAPIServer(responseJson string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, responseJson)
	}))
	slack.APIURL = fmt.Sprintf("%s/api/", ts.URL)
	return ts
}
