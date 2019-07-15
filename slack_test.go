package doumex

import (
	"testing"
)

func TestSlack_Send(t *testing.T) {
	s := Slack{
		WebHookURL: "https://hooks.slack.com/services/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",
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
