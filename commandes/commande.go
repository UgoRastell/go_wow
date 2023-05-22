package commande

import (
    "fmt"
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

    if event.ApplicationCommandData().Name == "delete" {
        guildID := event.GuildID
        commandID := "1100057933561221150"
    
        err := deleteCommandByID(s, guildID, commandID)
        if err != nil {
            fmt.Println("Error deleting command:", err)
        }
    }
    
}

