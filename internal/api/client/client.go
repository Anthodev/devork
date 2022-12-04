package client

import (
	"flag"
	"fmt"
	"github.com/anthodev/devork/internal/cmds"
	"github.com/anthodev/devork/internal/cmps/cmps_handlers"
	"github.com/anthodev/devork/internal/env"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
)

var s *discordgo.Session

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *env.BotToken)
	if err != nil {
		panic(err)
		//log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	presenceCommands        = cmds.BotCommands()
	presenceHandlers        = cmds.PresenceHandler()
	presenceButtonsHandlers = cmps_handlers.PresenceButtonsHandlers()
)

func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := presenceHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := presenceButtonsHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})
}

func RunDiscordApi() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()
	if err != nil {
		panic(err)
	}

	allCommands := append(presenceCommands)

	fmt.Println("Adding the commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(allCommands))

	for i, v := range allCommands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *env.GuildID, v)
		if err != nil {
			panic(err)
			//log.Fatalf("Error creating command '%v': %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer func(s *discordgo.Session) {
		err := s.Close()
		if err != nil {
			panic(err)
			//log.Fatalf("Error closing Discord session: %v", err)
		}
	}(s)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	fmt.Println("Bot is running. Press Ctrl+C to exit.")
	<-stop

	removeCommands(registeredCommands)
}

func removeCommands(registeredCommands []*discordgo.ApplicationCommand) {
	if *env.RemoveCmd {
		fmt.Println("Removing the commands...")
		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *env.GuildID, v.ID)
			if err != nil {
				fmt.Printf("Error removing command '%v': %v", v.Name, err)
			}
		}
	}
}
