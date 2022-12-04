package cmds

import (
	"github.com/anthodev/devork/internal/cmds/cmds_handlers"
	"github.com/bwmarrin/discordgo"
)

func PresenceHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"presence": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type != discordgo.InteractionApplicationCommand {
				return
			}

			cmds_handlers.CreatePresenceEmbedMessage(s, i, false)
		},
	}
}

func createPresenceInteractionResponse(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				cmds_handlers.CreatePresenceEmbed(),
			},
			Components: []discordgo.MessageComponent{
				cmds_handlers.CreatePresenceActionRow(),
			},
		},
	})

	if err != nil {
		panic(err)
	}
}
