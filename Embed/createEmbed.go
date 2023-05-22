package embed

import (

	"time"

	"github.com/bwmarrin/discordgo"

)

func CreateEmbed(authorName, authorIconURL, description string, fields []*discordgo.MessageEmbedField, imageURL, thumbnailURL, title string, color int) *discordgo.MessageEmbed {
	// Create the embed
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    authorName,
			IconURL: authorIconURL,
		},
		Color:       color,
		Description: description,
		Fields:      fields,
		Image: &discordgo.MessageEmbedImage{
			URL: imageURL,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: thumbnailURL,
		},
		Timestamp: time.Now().Format(time.RFC3339),
		Title:     title,
	}

	return embed
}

