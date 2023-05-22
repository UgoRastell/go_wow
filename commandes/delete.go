package commande

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"net/http"
)

func deleteCommandByID(s *discordgo.Session, guildID string, commandID string) error {
	// Call the deleteGuildCommand function with the session's application ID
	err := deleteGuildCommand(s, guildID, commandID)
	if err != nil {
		return err
	}

	return nil
}

func deleteGuildCommand(s *discordgo.Session, guildID string, commandID string) error {
	// Create the URL
	baseURL := "https://discord.com/api/v10"
	url := fmt.Sprintf("%s/applications/%s/guilds/%s/commands/%s/permissions", baseURL, s.State.User.ID, guildID, commandID)

	// Create the HTTP request
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	// Set the request headers
	req.Header.Set("Authorization", "Bearer <my_bearer_token>")

	// Send the HTTP request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete command permissions (status code: %d)", resp.StatusCode)
	}

	return nil
}
