package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	t "hackernewsbot/topics"
	u "hackernewsbot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var relevantTopics []string
var newStoriesIDs = "https://hacker-news.firebaseio.com/v0/newstories.json?print=pretty"
var newsInfos = "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"
var ids []int

//News struct is used to store the news
type News struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

//GetLatestNewsID returns the latest news id
func GetLatestNewsID() ([]int, error) {
	resp, err := http.Get(newStoriesIDs)
	if err != nil {
		return ids, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ids, fmt.Errorf("Error %s", resp.Status)
	}
	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		return ids, err
	}
	return ids, nil
}

// FetchNews returns the news titles and urls
func FetchNews(update tgbotapi.Update) {
	bot, err := u.Login()

	u.HandleError(err)

	response, err := GetLatestNewsID()
	u.HandleError(err)

	news := News{}

	for _, id := range response {
		resp, err := http.Get(fmt.Sprintf(newsInfos, id))
		u.HandleError(err)

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println(fmt.Sprintf("Error %s", resp.Status))
		}
		err = json.NewDecoder(resp.Body).Decode(&news)
		u.HandleError(err)

		topics, err := t.ReadTopics()
		u.HandleError(err)

		for _, topic := range topics {
			if news.Title != "" && news.URL != "" && strings.Contains(strings.ToLower(news.Title), topic) && !checkIfNewsIsInJSON(news) {
				addNewsToJSON(news)
				cleanJSONFile(news)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, news.Title+"\n"+news.URL)
				bot.Send(msg)
			}
		}
	}
}

func checkIfNewsIsInJSON(news News) bool {
	createFileIfNotExists()
	file, err := os.Open("news.json")
	u.HandleError(err)

	defer file.Close()

	var newsList []News
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&newsList)
	u.HandleError(err)

	for _, n := range newsList {
		if n.Title == news.Title {
			return true
		}
	}
	return false
}

func createFileIfNotExists() {
	_, err := os.Stat("news.json")
	if os.IsNotExist(err) {
		file, err := os.Create("news.json")
		u.HandleError(err)
		defer file.Close()
		encoder := json.NewEncoder(file)
		err = encoder.Encode([]News{})
		u.HandleError(err)
	}

}
func addNewsToJSON(news News) {
	file, err := os.Open("news.json")
	u.HandleError(err)

	decoder := json.NewDecoder(file)
	var newsArray []News
	err = decoder.Decode(&newsArray)
	u.HandleError(err)

	newsArray = append(newsArray, news)
	file, err = os.Create("news.json")
	u.HandleError(err)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(newsArray)
	u.HandleError(err)

	defer file.Close()
}

func cleanJSONFile(news News) {
	file, err := os.Open("news.json")
	u.HandleError(err)

	decoder := json.NewDecoder(file)
	var newsArray []News
	err = decoder.Decode(&newsArray)
	u.HandleError(err)

	if len(newsArray) > 500 {
		file, err = os.Create("news.json")
		u.HandleError(err)

		encoder := json.NewEncoder(file)
		err = encoder.Encode(newsArray[len(newsArray)-500:])
		u.HandleError(err)

		defer file.Close()
	}
}
