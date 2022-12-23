package main

import (
	"flag"
	"log"
	"os"

	"fukurokuju/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get the bot token from .env! Space for API tokens...
	botToken := os.Getenv("TOKEN")
	DEEPL_TOKEN := os.Getenv("DEEPL_TOKEN")
	// set GuID to flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	Gu := flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	// Start the bot
	bot.BotToken = botToken
	bot.DeeplToken = DEEPL_TOKEN
	bot.GuildID = *Gu
	bot.Run()

}
