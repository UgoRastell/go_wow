package commande

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
	"wow/Database"
	"wow/Embed"
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
	blizzardClientID = "4d50be5e687543d0a4754913047a8c3e"
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
	tokenBattleNet()
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
			Name:   "Instruction",
			Value:  "Veuillez cliquer sur le bouton pour vous authentifier (le bot ne récupère aucune information personnelle)",
			Inline: true,
		},
	}

	embedReponse := embed.CreateEmbed("GO_WOW", "https://play-lh.googleusercontent.com/PuPFgmLam2WNyul3lUQywQT5Y5sPgL6VzWSUAdXOS1oIQwHYnrB_MyfXCOrR4LzZcjeP=w240-h480-rw", "Permet de se connecter à son compte Battle.net",
		fields, "Authentification Battle.net", 28889)

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

	// Génère l'URL d'authentification OAuth2
	state := randomState(16)
	url := blizzardOauth2Config.AuthCodeURL(state, oauth2.AccessTypeOnline)

	button1 := embed.CreateButtonUrl("Login", url)

	// Crée la rangée d'actions
	actionRow := &discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{button1},
	}

	// Crée la réponse
	response := embed.ResponseEmbed(embedReponse, actionRow)

	err = s.InteractionRespond(i.Interaction, response)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi : ", err)
	}
}

func exchangeCodeForToken(code string) (*oauth2.Token, error) {
	blizzardClientSecret := token.BlizzardClientSecret()

	blizzardOauth2Config := &oauth2.Config{
		ClientID:     blizzardClientID,
		ClientSecret: blizzardClientSecret,
		RedirectURL: fmt.Sprintf("http://vps-e80a5a0d.vps.ovh.net/battle-net/"),
	}

	token, err := blizzardOauth2Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func tokenBattleNet() {
	token, err := exchangeCodeForToken("EUYBNCFIFSJHN9CWNQ9XMY88LFLFK5A9SP")
	if err != nil {
		fmt.Println("Erreur lors de l'échange du code d'autorisation : ", err)
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
		fmt.Println("Erreur lors de la conversion en JSON : ", err)
		return
	}

	fmt.Println(string(jsonResponse))
}





