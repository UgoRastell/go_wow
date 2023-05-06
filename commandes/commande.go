package commande

import (
    "github.com/bwmarrin/discordgo"
)

// onCommand est appelé lorsque la commande slash est exécutée
func onCommand(s *discordgo.Session, event *discordgo.InteractionCreate) {
    if event.ApplicationCommandData().Name == "example" {
        exemple(s, event)
    }

    if event.ApplicationCommandData().Name == "login" {
        oauth2LoginRegisterCommand(s, event)

    }
}

