package webhook

type TelegramPayload struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type Telegram struct {
	WebHookURL string           `json:"url"`
	Data       *TelegramPayload `json:"payload"`
}

func (t *Telegram) GetData() interface{} {
	return t.Data
}

func (t *Telegram) GetURL() string {
	return t.WebHookURL
}
