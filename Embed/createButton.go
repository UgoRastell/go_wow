package embed

import (
	"github.com/bwmarrin/discordgo"
)

func CreateButton(label string, style discordgo.ButtonStyle, customID string) *discordgo.Button {
	return &discordgo.Button{
		Label:    label,
		Style:    style,
		CustomID: customID,
	}
}

func CreateButtonUrl(label string, url string) *discordgo.Button {
	return &discordgo.Button{
		Label: label,
		Style: discordgo.LinkButton,
		URL:   url,
	}
}

func CreateSelectOption(label string, value string) *discordgo.SelectMenuOption {
	return &discordgo.SelectMenuOption{
		Label: label,
		Value: value,
	}
}

func CreateSelectMenu(customID string, options []*discordgo.SelectMenuOption) *discordgo.SelectMenu {
	convertedOptions := make([]discordgo.SelectMenuOption, len(options))
	for i, opt := range options {
		convertedOptions[i] = discordgo.SelectMenuOption{
			Label:       opt.Label,
			Value:       opt.Value,
			Description: opt.Description,
			Default:     opt.Default,
		}
	}

	return &discordgo.SelectMenu{
		CustomID: customID,
		Options:  convertedOptions,
	}
}
