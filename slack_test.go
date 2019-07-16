package doumex

import (
	"os"
	"testing"
)

func TestSlack_Send(t *testing.T) {
	url := os.Getenv("SLACK_TEST_URL")
	if url == "" {
		return
	}
	s := Slack{
		WebHookURL: url,
		Data: &SlackPayload{
			Text: "This is test code for slack.",
		},
	}

	if e := s.Send(); e != nil {
		for _, v := range e {
			t.Error(v)
		}
	}
}
