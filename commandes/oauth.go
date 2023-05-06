package commande

import (
    "github.com/bwmarrin/discordgo"
    "time"
)

func oauth2LoginRegisterCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
    // Create the embed
    embed := &discordgo.MessageEmbed{
        Author: &discordgo.MessageEmbedAuthor{
            Name: "My Bot",
            IconURL: "https://cdn.discordapp.com/avatars/123456789012345678/abcdefghijklmnopqrstuvwxy.jpg?size=256",
        },
        Color: 0x00ff00, // Green
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
        Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
        Title:     "I am an Embed",
    }

    // Create the button
    button := &discordgo.MessageComponent{
        Type: discordgo.ComponentActionRow,
        Components: []*discordgo.MessageButton{
            {
                Type:  discordgo.ButtonPrimary,
                Label: "Test",
                Style: discordgo.ButtonStylePrimary,
                CustomID: "test_button",
            },
        },
    }

    // Create the response
    response := &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content:   "Here is your embed!",
            Embeds:    []*discordgo.MessageEmbed{embed},
            Components: []discordgo.MessageComponent{*button},
        },
    }

    // Send the response
    err := s.InteractionRespond(i.Interaction, response)
    if err != nil {
        // Handle the error
    }
}



