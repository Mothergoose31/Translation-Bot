package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// ==========================================================================
// add whitelisted discord server ids here
var WhiteList = []string{}
	

// ==========================================================================

func checkWhitelist(id string) bool {
	for _, v := range WhiteList {
		if id == v {
			return true
		}
	}
	return false
}

// ==========================================================================

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			fmt.Println("DIDN'T ADD THE COMMANDS!")
		}
		registeredCommands[i] = cmd
	}
	// Set the playing status."
	err := s.UpdateListeningStatus("Christmas carols 🎅")
	if err != nil {
		fmt.Println("Error attempting to set my status")
	}
}

// ==========================================================================
var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "translate",
		Description: "Translate JP/EN | EN/JP",
		// String option here ->
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "query",
				Description: "Word or sentence.",
				Required:    true,
			},
		},
	},
	{
		Name:        "翻訳する",
		Description: "翻訳する JP/EN | EN/JP",
		// String option here ->
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "クエリ",
				Description: "単語や文章。",
				Required:    true,
			},
		},
	},
}

// ===============================================================================================
var CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"translate": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Get the options from the command data
		options := i.ApplicationCommandData().Options

		// Check if the "query" option is present
		var queryOption *discordgo.ApplicationCommandInteractionDataOption
		for _, opt := range options {
			if opt.Name == "query" {
				queryOption = opt
				break
			}
		}
		if queryOption == nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error: the \"query\" option was not provided.",
				},
			})
			return
		}
		//  if the guild is not in the whitelist, return an error

		// Get the query value and translate it
		fmt.Println(i.GuildID)
		if !checkWhitelist(i.GuildID) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error: This Discord Server is not whitelisted. Please contact the bot owner to get whitelisted. Goose#7218",
				},
			})
			return
		}

		query := queryOption.StringValue()
		output := TranslateToJapanese(query)

		// Send the response
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: output,
			},
		})
	},

	"翻訳する": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options

		var queryOption *discordgo.ApplicationCommandInteractionDataOption
		for _, option := range options {
			if option.Name == "クエリ" {
				queryOption = option
				break
			}
		}
		if queryOption == nil {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "エラー：「クエリ」オプションが提供されていま",
				},
			})
			return
		}
		// get the query value and translate it

		if !checkWhitelist(i.GuildID) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Error: This Discord Server is not whitelisted. Please contact the bot owner to get whitelisted. Goose#7218",
				},
			})
			return
		}

		query := queryOption.StringValue()
		output := TranslateToEnglish(query)

		// Send the response
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: output,
			},
		})
	},
}

// ==========================================================================
