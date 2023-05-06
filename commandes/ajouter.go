package commande

import (
	"github.com/bwmarrin/discordgo"
    "fmt"
)

func AddCommands(dg *discordgo.Session, guildID string) error {
    // Commande d'exemple 1
    cmd1 := &discordgo.ApplicationCommand{
        Name:        "example",
        Description: "Commande d'exemple",
    }
    _, err := RegisterCommand(dg, guildID, cmd1)
    if err != nil {
        return fmt.Errorf("Erreur lors de l'ajout de la commande slash 1 : %v", err)
    }
    
    // Commande d'exemple 2
    cmd2 := &discordgo.ApplicationCommand{
        Name:        "login",
        Description: "Ce login à son battle.net",
    }
    _, err = RegisterCommand(dg, guildID, cmd2)
    if err != nil {
        return fmt.Errorf("Erreur lors de l'ajout de la commande slash 1 : %v", err)
    }
    
    // Ajoute un gestionnaire pour l'événement de commande slash
    dg.AddHandler(onCommand)
    
    return nil
}


