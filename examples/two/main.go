package main

import (
	"encoding/json"
	"fmt"

	twitterscraper "github.com/firodj/twitter-scraper"
)

func main() {
	scraper := twitterscraper.New()
	tweet, err := scraper.GetTweet("1328684389388185600")
	if err != nil {
		panic(err)
	}
	j, err := json.MarshalIndent(tweet, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
