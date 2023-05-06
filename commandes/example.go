package commande

import (
	"github.com/bwmarrin/discordgo"
)		
		
func exemple(s *discordgo.Session, event *discordgo.InteractionCreate) {
    text := event.ApplicationCommandData().Options[0].StringValue()
    s.ChannelMessageSend(event.ChannelID, text)
}
