package catalpha

import (
	"github.com/bwmarrin/discordgo"
)

func Catalpha(m *discordgo.MessageCreate) string {
	msg := responsePrimal(m.Content)

	return msg
}
