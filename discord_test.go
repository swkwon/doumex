package doumex

import "testing"

func TestDiscord_Send(t *testing.T) {
	d := &Discord{
		WebHookURL: "https://discordapp.com/api/webhooks/598757949166452746/IjEBrzTau0up4Ao97eZ_ueQg2Mho31JbXthe9r2r7Dc0JCtmOTj1DPVEDDvcfMY_ybV0",
		Data: &DiscordPayload{
			Content: "hello world",
		},
	}
	if e := d.Send(); e != nil {
		for _, v := range e {
			t.Error(v)
		}
	}
}
