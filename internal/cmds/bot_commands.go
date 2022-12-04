package cmds

import (
	"github.com/anthodev/devork/internal/env"
	"github.com/bwmarrin/discordgo"
)

func BotCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "presence",
			Description: "Demander aux utilisateurs de mettre leur présence",
			GuildID:     *env.GuildID,
		},
	}
}

func ExampleCommands() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
		{
			Name:        "buttons",
			Description: "Test the buttons if you got courage",
		},
		{
			Name: "selects",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "multi",
					Description: "Multi-item select menu",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "single",
					Description: "Single-item select menu",
				},
			},
			Description: "Lo and behold: dropdowns are coming",
		},
	}
}
