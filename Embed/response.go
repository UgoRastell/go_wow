package embed

import (
	"github.com/bwmarrin/discordgo"
)

func ResponseEmbed(embedReponse *discordgo.MessageEmbed, actionRow *discordgo.ActionsRow) *discordgo.InteractionResponse {
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embedReponse},
			Components: []discordgo.MessageComponent{actionRow},
		},
	}

	return response
}
