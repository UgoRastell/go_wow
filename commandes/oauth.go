package commande

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	db "wow/Database"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

const (
	blizzardClientID     = "4d50be5e687543d0a4754913047a8c3e"
	blizzardClientSecret = "JztPwWCSK85RlcEfMm4MU74fQrmuiku7"
)

func randomState(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func oauth2LoginRegisterCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {

	client, err := db.ConnexionDatabase()
	if err != nil {
		fmt.Printf("Erreur lors de la connexion à la base de données : %v\n", err)
		return
	}

	// Vérifier si l'utilisateur existe déjà dans la base de données
	// Si l'utilisateur n'existe pas, l'ajouter à la base de données
	exists, err := checkUserExists(client, i.GuildID, i.Member.User.ID, i.Member.User.Username)
	if err != nil {
		fmt.Printf("Erreur lors de la vérification de l'existence de l'utilisateur : %v\n", err)
		return
	}
	if !exists {
		err = addNewUserToDatabase(client, i.GuildID, i.Member.User.ID, i.Member.User.Username)
		if err != nil {
			fmt.Printf("Erreur lors de l'insertion des données dans la collection : %v\n", err)
			return
		}
	}

	// Create the embed
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "GO_WOW",
			IconURL: "https://cdn.discordapp.com/avatars/123456789012345678/abcdefghijklmnopqrstuvwxy.jpg?size=256",
		},
		Color:       0x00ff00, // Green
		Description: "TEST !",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "I am a field",
				Value:  "I am a value",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "I am a second field",
				Value:  "I am a value",
				Inline: true,
			},
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://cdn.discordapp.com/avatars/123456789012345678/abcdefghijklmnopqrstuvwxy.jpg?size=256",
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: "https://cdn.discordapp.com/avatars/123456789012345678/abcdefghijklmnopqrstuvwxy.jpg?size=256",
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     "Connectez-vous à votre compte Battle.net!",
	}

	blizzardOauth2Config := &oauth2.Config{
		ClientID:     blizzardClientID,
		ClientSecret: blizzardClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://eu.battle.net/oauth/authorize",
			TokenURL: "https://eu.battle.net/oauth/token",
		},
		RedirectURL: "http://localhost/blizzard/",
		Scopes:      []string{"openid"},
	}

	// Generate the OAuth2 authentication URL
	state := randomState(16)
	url := blizzardOauth2Config.AuthCodeURL(state, oauth2.AccessTypeOnline)

	button1 := &discordgo.Button{
		Label: "Login",
		Style: discordgo.LinkButton,
		URL:   url,
	}

	// Create the action row
	actionRow := &discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button1},
	}

	// Create the response
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: []discordgo.MessageComponent{actionRow},
		},
	}

	err = s.InteractionRespond(i.Interaction, response)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi: ", err)
	}

}

func exchangeCodeForToken(code string) (*oauth2.Token, error) {
	blizzardOauth2Config := &oauth2.Config{
		ClientID:     blizzardClientID,
		ClientSecret: blizzardClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://eu.battle.net/oauth/authorize",
			TokenURL: "https://eu.battle.net/oauth/token",
		},
		RedirectURL: "http://localhost/blizzard/",
		Scopes:      []string{"openid"},
	}

	token, err := blizzardOauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func tokenBattleNet() ([]byte, error) {
	token, err := exchangeCodeForToken("EUELZT0AZNSKZH3UPG9YOQF5X5MEA5HGZ3")
	if err != nil {
		return nil, err
	}

	// Créer une structure avec les mêmes champs que la variable token
	response := TokenResponse{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		ExpiresIn:   int(token.Expiry.Sub(time.Now()).Seconds()),
	}

	// Encoder la structure en JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return jsonResponse, nil
}

func checkUserExists(client *mongo.Client, guildID string, userID string, username string) (bool, error) {
	// Get the collection from the database
	collection := client.Database("gowow").Collection("users")

	// Set the filter to search for the user
	filter := bson.M{"guild_id": guildID, "user_id": userID}

	// Set options to limit the result to one document
	options := options.FindOne().SetProjection(bson.M{"_id": 1})

	// Execute the query
	result := collection.FindOne(context.Background(), filter, options)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return false, nil
		}
		// An error occurred while executing the query
		return false, result.Err()
	}

	// User exists in the database
	return true, nil
}

func addNewUserToDatabase(client *mongo.Client, guildID string, userID string, username string) error {
	user := bson.M{
		"guild_id": guildID,
		"user_id":  userID,
		"username": username,
	}
	err := db.InsertDocument(client, "gowow", "users", user)
	if err != nil {
		fmt.Printf("Erreur lors de l'insertion des données dans la collection : %v\n", err)
		return err
	}
	return nil
}
