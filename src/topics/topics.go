package topics

import (
	"encoding/json"
	u "hackernewsbot/utils"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var topics []string

//ReadTopics reads the topics from the topics.json file
func ReadTopics() ([]string, error) {
	file, err := os.Open("topics.json")
	u.HandleError(err)
	defer file.Close()

	var topics []string
	err = json.NewDecoder(file).Decode(&topics)
	u.HandleError(err)
	return topics, nil
}

func convertTopicsToString(array []string) string {
	var topicString string

	for _, topic := range array {
		topicString = topicString + "-" + topic + "\n"
	}
	fullMessage := "Here are the topics you chose:\n" + topicString
	return fullMessage
}

//GetTopics returns the topics to the user
func GetTopics(update tgbotapi.Update) {
	bot, err := u.Login()
	u.HandleError(err)
	topics, err := ReadTopics()
	u.HandleError(err)
	var topicArray []string
	for _, topic := range topics {
		topicArray = append(topicArray, topic)
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, convertTopicsToString(topicArray))
	bot.Send(msg)
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Run /news to start recieving news every hour.")
	bot.Send(msg)
}
