package commande

import (
	"fmt"
	"math/rand"
	"time"
	db "wow/Database"
	embed "wow/Embed"
	"wow/tokens"
	"github.com/bwmarrin/discordgo"
	"golang.org/x/oauth2"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

const (
	blizzardClientID     = "4d50be5e687543d0a4754913047a8c3e"
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
	
	exists, err := db.CheckUserExists(client, i.GuildID, i.Member.User.ID, i.Member.User.Username)
	if err != nil {
		fmt.Printf("Erreur lors de la vérification de l'existence de l'utilisateur : %v\n", err)
		return
	}
	if !exists {
		err = db.AddNewUserToDatabase(client, i.GuildID, i.Member.User.ID, i.Member.User.Username)
		if err != nil {
			fmt.Printf("Erreur lors de l'insertion des données dans la collection : %v\n", err)
			return
		}
	}

	fields := []*discordgo.MessageEmbedField{
		&discordgo.MessageEmbedField{
			Name:   "Custom Field 1",
			Value:  "Custom Value 1",
			Inline: true,
		},
	}
	
	embedReponse := embed.CreateEmbed("Custom Author", "https://picsum.photos/200", "Custom Description", 
	fields, "https://example.com/image.jpg", "https://example.com/thumbnail.jpg", "Custom Title", 0xff0000)

	blizzardClientSecret := token.BlizzardClientSecret()

	blizzardOauth2Config := &oauth2.Config{
		ClientID:     blizzardClientID,
		ClientSecret: blizzardClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://eu.battle.net/oauth/authorize",
			TokenURL: "https://eu.battle.net/oauth/token",
		},
		RedirectURL: fmt.Sprintf("http://vps-e80a5a0d.vps.ovh.net/battle-net/?user_id=%s", i.Member.User.ID),
		Scopes:      []string{"openid"},
	}

	// Generate the OAuth2 authentication URL
	state := randomState(16)
	url := blizzardOauth2Config.AuthCodeURL(state, oauth2.AccessTypeOnline)

	button1 := embed.CreateButtonUrl("Login", url)

	// Create the action row
	actionRow := &discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button1},
	}

	// Create the response
	response := embed.ResponseEmbed(embedReponse , actionRow)

	err = s.InteractionRespond(i.Interaction, response)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi: ", err)
	}

}





