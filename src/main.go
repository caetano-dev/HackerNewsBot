package main

import (
	"log"
	"time"

	a "hackernewsbot/api"
	t "hackernewsbot/topics"
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
				t.GetTopics(update)
			}

			if update.Message.Text == "/news" {
				for {
					a.FetchNews(update)
					time.Sleep(time.Hour)
				}
			}
		}

	}
}
