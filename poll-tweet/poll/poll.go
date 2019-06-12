package main

import (
	"errors"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	"github.com/golang/glog"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
)

// SearchResultCount controls the batch size of the search results
var SearchResultCount = 100

// TwitterTimeLayout is the format Twitter uses for CreatedAt
var TwitterTimeLayout = "Mon Jan 2 15:04:05 -0700 2006"

func findMaxSinceID(tweets []twitter.Tweet, prevSinceID int64) int64 {
	maxSinceID := prevSinceID
	for _, tweet := range tweets {
		if tweet.ID > prevSinceID {
			maxSinceID = tweet.ID
		}
	}
	return maxSinceID
}

// PollAllTweets loops over all search specs and polls tweets later than Since_ID
func PollAllTweets(s services.StateStorer) ([]twitter.Tweet, error) {
	glog.Info("PollAllTweets started...")
	state, err := s.GetState()
	if err != nil {
		return nil, err
	}
	var allTweets []twitter.Tweet
	for query, sinceID := range state {
		tweets, err := PollTweets(query, sinceID)
		if err != nil {
			return nil, err
		}
		allTweets = append(allTweets, tweets...)
		state[query] = findMaxSinceID(tweets, sinceID)
	}
	if err := s.SetState(state); err != nil {
		return nil, err
	}
	glog.Info("PollAllTweets finished...")
	return allTweets, nil
}

// PollTweets is polling tweets from Twitter
func PollTweets(query string, sinceID int64) ([]twitter.Tweet, error) {
	glog.Infof("PollTweets for '%s' started...", query)
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		return nil, errors.New("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	searchParams := &twitter.SearchTweetParams{
		Query:           query,
		Lang:            "en",
		ResultType:      "popular",
		Count:           SearchResultCount,
		SinceID:         sinceID,
		IncludeEntities: twitter.Bool(true),
	}
	search, _, err := client.Search.Tweets(searchParams)
	if err != nil {
		return nil, err
	}
	for i := range search.Statuses {
		search.Statuses[i].Source = query
	}
	glog.Infof("PollTweets for '%s' finished, found %d tweets...", query, len(search.Statuses))
	return search.Statuses, nil
}

// PublishTweets sends tweets via event publisher
func PublishTweets(ep services.EventPublisher, tweets []twitter.Tweet) error {
	glog.Info("PublishTweets started...")
	var events []services.Event
	for _, tweet := range tweets {
		tm, err := time.Parse(TwitterTimeLayout, tweet.CreatedAt)
		if err != nil {
			return err
		}
		events = append(events, services.Event{
			ID:        string(tweet.ID),
			Shard:     tweet.Source,
			Timestamp: tm,
			Source:    "Poll-Tweet",
			EventType: "Tweet",
			Payload:   tweet,
		})
	}
	if err := ep.PublishEvents(events); err != nil {
		return err
	}
	glog.Info("PublishTweets finished...")
	return nil
}
