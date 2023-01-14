package main

import (
    "context"
    "fmt"
    twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
    scraper := twitterscraper.New().SetSearchMode(twitterscraper.SearchUsers)
    for profile := range scraper.SearchProfiles(context.Background(), "Twitter", 50) {
        if profile.Error != nil {
            panic(profile.Error)
        }
        fmt.Println(profile.Name)
    }
}
