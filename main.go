package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	data "wow/Mountdata"
	token "wow/tokens"

	"github.com/bwmarrin/discordgo"
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
	data.Data()
	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	dg.Close()
}
