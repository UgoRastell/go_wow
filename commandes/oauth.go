package commande

import (
	"fmt"
	"time"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/oauth2"
	"math/rand"
	"context"
    "encoding/json"
)

type TokenResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
}

const (
	// These values are obtained by creating a new OAuth2 application on the Blizzard Developer Portal:
	// https://develop.battle.net/access/clients
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

    err := s.InteractionRespond(i.Interaction, response)
    if err != nil {
        fmt.Println("Erreur lors de l'envoi: ", err)
    }

	tokenBattleNet()

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

func tokenBattleNet() {
    token, err := exchangeCodeForToken("EUELZT0AZNSKZH3UPG9YOQF5X5MEA5HGZ3")
    if err != nil {
        return
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
        // gérer l'erreur
    }

    fmt.Println(string(jsonResponse))
    return
}










