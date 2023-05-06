package commande

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func RegisterCommand(dg *discordgo.Session, guildID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
    // Enregistre la commande slash
    createdCmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, cmd)
    if err != nil {
        return nil, fmt.Errorf("Erreur lors de la création de la commande slash : %v", err)
    }
    
    return createdCmd, nil
}










