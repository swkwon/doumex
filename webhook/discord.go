package webhook

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type DiscordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type URL struct {
	URL string `json:"url"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type Embed struct {
	Author      *Author        `json:"author"`
	Title       string         `json:"title"`
	URL         string         `json:"url"`
	Description string         `json:"description"`
	Color       int32          `json:"color"`
	Fields      []DiscordField `json:"fields"`
	Thumbnail   URL            `json:"thumbnail"`
	Image       URL            `json:"image"`
	Foot        Footer         `json:"footer"`
}

type DiscordPayload struct {
	UserName  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	Content   string  `json:"content,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

type Discord struct {
	WebHookURL string          `json:"url"`
	Data       *DiscordPayload `json:"payload"`
}

func (d *Discord) GetData() interface{} {
	return d.Data
}

func (d *Discord) GetURL() string {
	return d.WebHookURL
}
