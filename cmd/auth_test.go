package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nlopes/slack"
)

func TestSubCmdAuthTest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
    "ok": true,
    "url": "https:\/\/example.slack.com\/",
    "team": "example",
    "user": "example",
    "team_id": "T3A1K77EH",
    "user_id": "U39SENP7D"
}`)
	}))
	defer ts.Close()

	slack.APIURL = fmt.Sprintf("%s/api/", ts.URL)

	subCmdAuthTest := SubCmdAuthTest()
	stdout, stderr, err := executeCommand(subCmdAuthTest, "auth", "test")
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}

	expected := fmt.Sprintf("user: example\nurl: https://example.slack.com/\n")

	if stdout != expected {
		t.Errorf("got stdout and expected output is not equal\ngot: %s\nexpected: %s\n", stdout, expected)
	}

	if len(stderr) != 0 {
		t.Errorf("expect stderr is nothing but got:\n%s", stderr)
	}
}
