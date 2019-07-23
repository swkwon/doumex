[![Build Status](https://travis-ci.org/swkwon/doumex.svg?branch=master)](https://travis-ci.org/swkwon/doumex)
# doumex

Doumex makes it easy to use the incoming webhook features of Slack, Telegram, Discord, and Zulip.

## Discord
```go
discord := &doumex.Discord{
	WebHookURL: {WEBHOOK_URL},
	Data: &doumex.DiscordPayload {
		UserName: "my bot",
		Content: "Hello world",
	},
}

discord.Send()
```

## Slack
```go
slack := &doumex.Slack{
	WebHookURL: {WEBHOOK_URL},
	Data: &doumex.SlackPayload {
		Text: "Hello world",
	},
}

slack.Send()
```

## Telegram
```go
telegram := &doumex.Telegram{
    WebHookURL: {WEBHOOK_URL},
    Data: &TelegramPayload{
        ChatID: {CHAT_ID},
        Text: "Hello world",
    },
}

telegram.Send()
```

## Zulip
```go
zulip := &doumex.Zulip{
    WebHookURL: {WEBHOOK_URL},
    Bot:        {BOT_NAME},
    APIKey:     {BOT_API_KEY},
    Data: &ZulipData{
        Type:    {stream | private},
        To:      {STREAM_NAME},
        Subject: {TOPIC_NAME},
        Content: "Hello world",
    },
}

zulip.Send()
```