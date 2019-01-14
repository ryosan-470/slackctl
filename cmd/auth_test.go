package cmd

import (
	"fmt"
	"testing"
)

func TestSubCmdAuthTest(t *testing.T) {
	testCases := []struct {
		jsonTxt  string
		expected string
		err      string
	}{
		{
			jsonTxt: `{
    "ok": true,
    "url": "https:\/\/example.slack.com\/",
    "team": "example",
    "user": "example",
    "team_id": "T3A1K77EH",
    "user_id": "U39SENP7D"
}`,
			expected: fmt.Sprintf("user: example\nurl: https://example.slack.com/\n"),
			err:      "",
		},
		{
			jsonTxt: `{
    "ok": false,
    "error": "invalid_auth"
}`,
			expected: "",
			err:      "Calling auth.test is failed: invalid_auth",
		},
	}

	subCmdAuthTest := SubCmdAuthTest()
	for i, testCase := range testCases {
		ts := dummySlackAPIServer(testCase.jsonTxt)
		defer ts.Close()

		stdout, stderr, err := executeCommand(subCmdAuthTest, "auth", "test")
		if err != nil && testCase.err != err.Error() {
			t.Errorf("got: %v, but want %v", err, testCase.err)
		}

		if stdout != testCase.expected {
			t.Errorf("testCases[%d] got stdout and expected output is not equal\ngot: %s\nexpected: %s\n", i, stdout, testCase.expected)
		}

		if len(stderr) != 0 {
			t.Errorf("expect stderr is nothing but got:\n%s", stderr)
		}
	}
}
