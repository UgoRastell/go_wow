package button

import (
	"github.com/bwmarrin/discordgo"
)

func HandleButtonClick(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.MessageComponentData().CustomID == "loginButton" {
		loginBN(s, i)
	}
	
}
