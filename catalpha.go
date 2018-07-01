package catalpha

import "github.com/bwmarrin/discordgo"

func Catalpha(m *discordgo.MessageCreate) string {
	return simpleResponse(m.Content)
}
