package doumex

import (
	"os"
	"testing"
)

func TestTelegram_Send(t *testing.T) {
	url := os.Getenv("TG_TEST_URL")
	chatID := os.Getenv("TG_CHAT_ID")

	if url == "" || chatID == "" {
		return
	}

	tg := &Telegram{
		WebHookURL: url,
		Data: &TelegramPayload{
			ChatID: chatID,
			Text: "This is test code for telegram.",
		},
	}

	if e := tg.Send(); e != nil {
		for _, v := range e {
			t.Error(v)
		}
	}
}
