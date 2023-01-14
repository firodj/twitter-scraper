package twitterscraper

import "time"

type (
	// Video type.
	Video struct {
		ID      string	`json:"id"`
		Preview string	`json:"preview"`
		URL     string	`json:"url"`
	}

	// Tweet type.
	Tweet struct {
		Hashtags         []string `json:"hashtags"`
		HTML             string  `json:"html"`
		ID               string  `json:"id"`
		InReplyToStatus  *Tweet  `json:"in_reply_to_status"`
		IsQuoted         bool	`json:"is_quoted"`
		IsPin            bool	`json:"is_pin"`
		IsReply          bool	`json:"is_reply"`
		IsRetweet        bool	`json:"is_retweet"`
		Likes            int	`json:"likes"`
		PermanentURL     string	`json:"permanent_url"`
		Photos           []string	`json:"photos"`
		Place            *Place		`json:"place"`
		QuotedStatus     *Tweet	`json:"quoted_status"`
		Replies          int	`json:"replies"`
		Retweets         int	`json:"retweets"`
		RetweetedStatus  *Tweet	`json:"retweeted_status"`
		Text             string	`json:"text"`
		TimeParsed       time.Time	`json:"time_parsed"`
		Timestamp        int64	`json:"timestamp"`
		URLs             []string `json:"urls"`
		UserID           string	`json:"user_id"`
		Username         string	`json:"username"`
		Videos           []Video	`json:"videos"`
		SensitiveContent bool	`json:"sensitive_content"`
	}

	// ProfileResult of scrapping.
	ProfileResult struct {
		Profile
		Error error	`json:"error"`
	}

	// TweetResult of scrapping.
	TweetResult struct {
		Tweet
		Error error `json:"error"`
	}

	legacyUser struct {
		CreatedAt   string `json:"created_at"`
		Description string `json:"description"`
		Entities    struct {
			URL struct {
				Urls []struct {
					ExpandedURL string `json:"expanded_url"`
				} `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FavouritesCount      int      `json:"favourites_count"`
		FollowersCount       int      `json:"followers_count"`
		FriendsCount         int      `json:"friends_count"`
		IDStr                string   `json:"id_str"`
		ListedCount          int      `json:"listed_count"`
		Name                 string   `json:"name"`
		Location             string   `json:"location"`
		PinnedTweetIdsStr    []string `json:"pinned_tweet_ids_str"`
		ProfileBannerURL     string   `json:"profile_banner_url"`
		ProfileImageURLHTTPS string   `json:"profile_image_url_https"`
		Protected            bool     `json:"protected"`
		ScreenName           string   `json:"screen_name"`
		StatusesCount        int      `json:"statuses_count"`
		Verified             bool     `json:"verified"`
	}

	Place struct {
		ID          string `json:"id"`
		PlaceType   string `json:"place_type"`
		Name        string `json:"name"`
		FullName    string `json:"full_name"`
		CountryCode string `json:"country_code"`
		Country     string `json:"country"`
		BoundingBox struct {
			Type        string        `json:"type"`
			Coordinates [][][]float64 `json:"coordinates"`
		} `json:"bounding_box"`
	}

	fetchProfileFunc func(query string, maxProfilesNbr int, cursor string) ([]*Profile, string, error)
	fetchTweetFunc   func(query string, maxTweetsNbr int, cursor string) ([]*Tweet, string, error)
)
