package webhook

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
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

func (z *Zulip) Send() error {
	req, e := http.NewRequest("POST", z.WebHookURL, bytes.NewBuffer([]byte(z.makeContent())))
	if e != nil {
		return e
	}
	req.SetBasicAuth(z.Bot, z.APIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return errors.New(res.Status)
	}
	return nil
}
