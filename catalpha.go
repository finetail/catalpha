package catalpha

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func Catalpha(m *discordgo.MessageCreate) string {
	msg := simpleResponse(m.Content)
	if msg == "" {
		if time.Now().Unix()%59 == 0 {
			msg = "ヤバイですね！"
		}
	}
	return msg
}
