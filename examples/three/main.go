package main

import (
	"context"
	"encoding/json"
	"fmt"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	scraper := twitterscraper.New()
	i := 0
	for tweet := range scraper.SearchTweets(context.Background(),
		"twitter scraper data -filter:retweets", 60) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		j, err := json.MarshalIndent(tweet, "", "  ")
		if err != nil {
			panic(err)
		}
		i++
		fmt.Println(i)
		fmt.Println(string(j))
	}
}
