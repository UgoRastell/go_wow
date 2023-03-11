package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"


	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"wow/commande/commandes"
)

func main() {
	// Charger les variables d'environnement depuis le fichier .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
		return
	}

	// Récupérer le token de bot depuis les variables d'environnement
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		fmt.Println("BOT_TOKEN environment variable not found.")
		return
	}

	// Créer une nouvelle session DiscordGo en utilisant le token de bot
	dg, err := discordgo.New("Bot " + token)
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



