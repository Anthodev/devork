package env

import (
	"flag"
	"os"
)

var (
	GuildID   = flag.String("guild", os.Getenv("GUILD_ID"), "Test guild ID. If not passed - bot registers cmds globally")
	BotToken  = flag.String("token", os.Getenv("BOT_TOKEN"), "Bot access token")
	RemoveCmd = flag.Bool("remove", true, "Remove all cmds after shutdowning or not")
)
