package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken   string
	DeeplToken string
	GuildID    string
)

// ===============

func Run() {
	fmt.Println("Got Keys: ", BotToken)
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the handlers
	dg.AddHandler(Ready)
	// dg.AddHandler(messageCreate)
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// Add the intents for the bot!
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open a websocket connection to Discord and begin listening.
	dg.Debug = true
	err = dg.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
