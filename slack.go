package doumex

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type SlackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Action struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Url   string `json:"url"`
	Style string `json:"style"`
}

type Attachment struct {
	Fallback     *string       `json:"fallback"`
	Color        *string       `json:"color"`
	PreText      *string       `json:"pretext"`
	AuthorName   *string       `json:"author_name"`
	AuthorLink   *string       `json:"author_link"`
	AuthorIcon   *string       `json:"author_icon"`
	Title        *string       `json:"title"`
	TitleLink    *string       `json:"title_link"`
	Text         *string       `json:"text"`
	ImageUrl     *string       `json:"image_url"`
	Fields       []*SlackField `json:"fields"`
	Footer       *string       `json:"footer"`
	FooterIcon   *string       `json:"footer_icon"`
	Timestamp    *int64        `json:"ts"`
	MarkdownIn   *[]string     `json:"mrkdwn_in"`
	Actions      []*Action     `json:"actions"`
	CallbackID   *string       `json:"callback_id"`
	ThumbnailUrl *string       `json:"thumb_url"`
}

type SlackPayload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

type Slack struct {
	WebHookURL string        `json:"url"`
	Data       *SlackPayload `json:"payload"`
}

func (s *Slack) Send() []error {
	jsonData, err := json.Marshal(s.Data)
	if err != nil {
		return []error{err}
	}
	request := gorequest.New()
	resp, _, errs := request.Post(s.WebHookURL).Send(string(jsonData)).End()
	if errs != nil {
		return errs
	}
	if resp.StatusCode >= 400 {
		return []error{fmt.Errorf("error sending msg. status: %v", resp.StatusCode)}
	}
	return nil
}
