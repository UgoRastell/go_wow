package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	commande "wow/Commandes"
	"github.com/bwmarrin/discordgo"

	token "wow/tokens"
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

	

	// Ajoute les commandes slash
	err = commande.AddCommands(dg, "1083387976588984490")
	if err != nil {
		fmt.Println("Erreur lors de l'ajout des commandes : ", err)
		return
	}

	token.UpdateAccessToken()
	
	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	dg.Close()
}
