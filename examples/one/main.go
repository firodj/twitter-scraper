package main

import (
	"context"
	"encoding/json"
	"fmt"

	twitterscraper "github.com/firodj/twitter-scraper"
)

func main() {
	scraper := twitterscraper.New()

	for tweet := range scraper.GetTweets(context.Background(), "Twitter", 10) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		j, err := json.MarshalIndent(tweet, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(j))
	}
}
