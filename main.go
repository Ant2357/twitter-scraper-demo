package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

type Tweet struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func main() {
	scraper := twitterscraper.New()

	tweets := make([]Tweet, 0)
	for tweet := range scraper.GetTweets(context.Background(), "nhk_news", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}

		fmt.Println(tweet.Text)
		tweets = append(tweets, Tweet{
			Id:   tweet.ID,
			Text: tweet.Text,
		})
	}

	file, err := json.MarshalIndent(tweets, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("tweets.json", []byte(file), 0644)
}
