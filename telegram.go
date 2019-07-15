package doumex

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type TelegramPayload struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type Telegram struct {
	WebHookURL string           `json:"url"`
	Data       *TelegramPayload `json:"payload"`
}

func (t *Telegram) Send() []error {
	jsonData, err := json.Marshal(t.Data)
	if err != nil {
		return nil
	}
	request := gorequest.New()
	resp, _, errs := request.Post(t.WebHookURL).Send(string(jsonData)).End()
	if errs != nil {
		return errs
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("error sending msg. status: %v", resp.StatusCode)}
	}
	return nil
}
