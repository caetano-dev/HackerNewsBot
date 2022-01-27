package utils

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

//Login will initiate the bot and return the bot object
func Login() (*tgbotapi.BotAPI, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	HandleError(err)
	return bot, nil
}
