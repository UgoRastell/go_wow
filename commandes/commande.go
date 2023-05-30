package commande

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// onCommand est appelé lorsque la commande slash est exécutée
func onCommand(s *discordgo.Session, event *discordgo.InteractionCreate) {
	if event.ApplicationCommandData().Name == "login" {
		oauth2LoginRegisterCommand(s, event)
	}

	if event.ApplicationCommandData().Name == "reset_monture" {
		Run(s, event)
	}

}

func AddCommands(dg *discordgo.Session, guildID string) error {

	cmd1 := &discordgo.ApplicationCommand{
		Name:        "example",
		Description: "Commande d'exemple",
	}
	_, err := RegisterCommand(dg, guildID, cmd1)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'ajout de la commande slash 1 : %v", err)
	}

	cmd2 := &discordgo.ApplicationCommand{
		Name:        "login",
		Description: "Ce login à son battle.net",
	}
	_, err = RegisterCommand(dg, guildID, cmd2)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'ajout de la commande slash 2 : %v", err)
	}

	cmd3 := &discordgo.ApplicationCommand{
		Name:        "reset_monture",
		Description: "Reset la data base des montures wow",
	}
	_, err = RegisterCommand(dg, guildID, cmd3)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'ajout de la commande slash 3 : %v", err)
	}

	// Ajoute un gestionnaire pour l'événement de commande slash
	dg.AddHandler(onCommand)

	return nil
}

func RegisterCommand(dg *discordgo.Session, guildID string, cmd *discordgo.ApplicationCommand) (*discordgo.ApplicationCommand, error) {
	// Enregistre la commande slash
	createdCmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, cmd)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la création de la commande slash : %v", err)
	}

	return createdCmd, nil
}
