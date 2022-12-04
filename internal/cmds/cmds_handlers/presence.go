package cmds_handlers

import (
	"fmt"
	"github.com/anthodev/devork/internal/api/response"
	"github.com/anthodev/devork/internal/cmds/cmds_struct"
	"github.com/bwmarrin/discordgo"
	"regexp"
	"strconv"
	"time"
)

var (
	startDate = time.Now()
	endDate   = time.Now()
)

func CreatePresenceEmbedMessage(s *discordgo.Session, i *discordgo.InteractionCreate, isUpdated bool) {
	options := i.ApplicationCommandData().Options
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

	if len(options) > 0 {
		for _, option := range options {
			optionsMap[option.Name] = option
		}

		isUpdated, _ = strconv.ParseBool(optionsMap["update"].StringValue())
	}

	originalDate := time.Now()
	currentDate := time.Now()

	if isUpdated {
		for _, v := range i.Interaction.Message.Embeds {
			if v.Type == "rich" {
				title := v.Title

				re := regexp.MustCompile(`\d+`)
				match := re.FindSubmatch([]byte(title))

				if strconv.Itoa(weekStartDate(originalDate).Day()) == string(match[0]) {
					startDate = weekStartDate(currentDate)
					endDate = weekEndDate(currentDate)
				} else {
					startDate = nextWeekStartDate(currentDate)
					endDate = nextWeekEndDate(currentDate)
				}

				CreatePresenceMessage(s, i, startDate, endDate)
			}
		}
	} else {
		if int(originalDate.Weekday()) == 0 || int(originalDate.Weekday()) >= 2 {
			startDate = nextWeekStartDate(currentDate)
			endDate = nextWeekEndDate(currentDate)
		} else {
			startDate = weekStartDate(currentDate)
			endDate = weekEndDate(currentDate)
		}

		CreatePresenceMessage(s, i, startDate, endDate)
	}
}

func CreatePresenceMessage(s *discordgo.Session, i *discordgo.InteractionCreate, startDate time.Time, endDate time.Time) {
	embed := CreatePresenceEmbed()

	actionRow := CreatePresenceActionRow()

	response.SendEmbed(embed, actionRow, s, i)
}

func CreatePresenceEmbed() *discordgo.MessageEmbed {
	embed := cmds_struct.NewGenericEmbed(
		fmt.Sprintf("Présence au bureau pour la semaine du %s au %s", startDate.Format("02/01"), endDate.Format("02/01")),
		"Indiquer vos jours de présence prévus au bureau pour la semaine indiquée.",
		0x004daa,
	)

	fields := []*discordgo.MessageEmbedField{
		{
			"**Lundi**",
			"---",
			false,
		},
		{
			"**Mardi**",
			"---",
			false,
		},
		{
			"**Mercredi**",
			"---",
			false,
		},
		{
			"**Jeudi**",
			"---",
			false,
		},
		{
			"**Vendredi**",
			"---",
			false,
		},
	}

	embed.Fields = fields

	return embed
}

func CreatePresenceActionRow() discordgo.ActionsRow {
	actionRow := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			cmds_struct.NewButton().
				SetLabel("Lundi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d1"),
			cmds_struct.NewButton().
				SetLabel("Mardi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d2"),
			cmds_struct.NewButton().
				SetLabel("Mercredi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d3"),
			cmds_struct.NewButton().
				SetLabel("Jeudi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d4"),
			cmds_struct.NewButton().
				SetLabel("Vendredi").
				SetStyle(discordgo.PrimaryButton).
				SetCustomID("d5"),
		},
	}

	return actionRow
}

//func UpdatePresenceMessage(s *discordgo.Session, i *discordgo.InteractionCreate) discordgo.InteractionResponseData {
//	for _, v = range i.Interaction.Message.Embeds {
//		if v.Type == discordgo.EmbedTypeRich {
//			for _, f = range v.Fields {
//				users := strings.Split(f.Value, ",")
//
//				for ind, u = range users {
//					if u == "---" {
//						removeUserFromSlice(users, ind)
//					}
//				}
//
//				if f.Value == "**Lundi**" {
//					f.Value = "---"
//				} else if f.Value == "**Mardi**" {
//					f.Value = "---"
//				} else if f.Value == "**Mercredi**" {
//					f.Value = "---"
//				} else if f.Value == "**Jeudi**" {
//					f.Value = "---"
//				} else if f.Value == "**Vendredi**" {
//					f.Value = "---"
//				}
//			}
//			title := v.Title
//
//			re := regexp.MustCompile(`\d+`)
//			match := re.FindSubmatch(title)
//
//			if strconv.Itoa(weekStartDate(originalDate).Day()) == string(match[0]) {
//				startDate = weekStartDate(currentDate)
//				endDate = weekEndDate(currentDate)
//			} else {
//				startDate = nextWeekStartDate(currentDate)
//				endDate = nextWeekEndDate(currentDate)
//			}
//		}
//	}
//}

func weekStartDate(date time.Time) time.Time {
	offset := (int(time.Monday) - int(date.Weekday()) - 7) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)

	return result
}

func weekEndDate(date time.Time) time.Time {
	offset := (int(time.Friday) - int(date.Weekday()) - 1) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)

	return result
}

func nextWeekStartDate(date time.Time) time.Time {
	offset := (int(time.Monday) - int(date.Weekday()) + 7) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)

	return result
}

func nextWeekEndDate(date time.Time) time.Time {
	offset := (int(time.Friday) - int(date.Weekday()) + 7) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)

	return result
}
func removeUserFromSlice(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	slice[len(slice)-1] = ""
	slice = slice[:len(slice)-1]

	return slice
}
