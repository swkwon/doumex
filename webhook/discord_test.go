package webhook

import (
	"os"
	"testing"
)

func TestDiscord_Send(t *testing.T) {
	url := os.Getenv("DISCORD_TEST_URL")
	if url == "" {
		return
	}
	d := &Discord{
		WebHookURL: url,
		Data: &DiscordPayload{
			Content: "This is test code for discord.",
		},
	}
	if e := Send(d); e != nil {
		t.Error(e)
	}
}
