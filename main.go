package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"


	"github.com/bwmarrin/discordgo"
	"wow/commandes"
	"wow/tokens"
)

func main() {

	// Créer une nouvelle session DiscordGo en utilisant le token de bot
	dg, err := discordgo.New("Bot " + token.BotToken())
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Se connecter à Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error connecting to Discord: ", err)
		return
	}
	
	

	// Attendre un signal pour fermer le bot
	fmt.Println("Bot is now running.")
	commande.MountIndex()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Fermer la session DiscordGo
	dg.Close()
}



