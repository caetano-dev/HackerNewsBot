package main

import (
	"fmt"
	"log"
	"time"

	a "hackernewsbot/api"
	u "hackernewsbot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, error := u.Login()
	u.HandleError(error)
	log.Printf("Starting %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				for {
					a.FetchNews(update)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "end")
					bot.Send(msg)
					fmt.Println("I sleep")
					time.Sleep(time.Hour)
				}
			}

		}

	}
}