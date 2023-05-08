package button

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func loginBN(s *discordgo.Session, i *discordgo.InteractionCreate) {
	
	// Get the channel ID from the interaction
	channelID := i.ChannelID

	// Create the message
	message := "Hello, you pressed the login button!"

	// Send the message to the channel
	_, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		fmt.Println("Error sending message: ", err)
		return
	}
}
