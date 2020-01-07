package doumex

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

type ZulipPayload struct {
	Type    string `json:"type"`
	To      string `json:"to"`
	Subject string `json:"subject,omitempty"`
	Content string `json:"content"`
}

type Zulip struct {
	WebHookURL string        `json:"url"`
	Data       *ZulipPayload `json:"payload"`
	Bot        string        `json:"bot"`
	APIKey     string        `json:"apikey"`
}

func (z *Zulip) makeContent() string {
	return fmt.Sprintf(`type=%s&to=%s&subject=%s&content=%s`, z.Data.Type, z.Data.To, z.Data.Subject, z.Data.Content)
}

func (z *Zulip) Send() []error {
	request := gorequest.New()
	request.SetBasicAuth(z.Bot, z.APIKey)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	content := z.makeContent()
	resp, _, errs := request.Post(z.WebHookURL).Send(content).End()
	if errs != nil {
		return errs
	}

	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("error sending msg. status: %v", resp.StatusCode)}
	}
	return nil
}
