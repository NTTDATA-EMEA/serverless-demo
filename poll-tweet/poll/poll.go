package main

import (
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/okoeth/serverless-demo/commons/pkg/services"

	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// SearchResultCount controls the batch size of the search results
var SearchResultCount = 500

// TwitterTimeLayout is the format Twitter uses for CreatedAt
var TwitterTimeLayout = "Mon Jan 2 15:04:05 -0700 2006"

type TwitterSearchResults struct {
	Query   string
	SinceId int64
	Tweets  []twitter.Tweet
}

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{})
}

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
	log.Info("PollAllTweets started...")
	state, err := s.GetState()
	if err != nil {
		return nil, err
	}
	// run queries concurrently
	resc, errc := make(chan *TwitterSearchResults), make(chan error)
	for query, sinceID := range state {
		go func(q string, s int64) {
			res, err := PollTweets(q, s)
			if err != nil {
				errc <- fmt.Errorf("error for query %s: %v", q, err)
				return
			}
			resc <- res
		}(query, sinceID)
	}
	// collect results
	var allTweets []twitter.Tweet
	queryTimeout, err := strconv.Atoi(os.Getenv("TWITTER_QUERY_TIMEOUT_SEC"))
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(state); i++ {
		select {
		case res := <-resc:
			allTweets = append(allTweets, res.Tweets...)
			state[res.Query] = findMaxSinceID(res.Tweets, res.SinceId)
		case err := <-errc:
			return nil, err
		case <-time.After(time.Duration(queryTimeout) * time.Second):
			return nil, errors.New("polling timeout, Twitter query took to long")
		}
	}
	if err := s.SetState(state); err != nil {
		return nil, err
	}
	log.Info("PollAllTweets finished...")
	return allTweets, nil
}

// PollTweets is polling tweets from Twitter
func PollTweets(query string, sinceID int64) (*TwitterSearchResults, error) {
	log.WithFields(log.Fields{
		"hashtag": query,
	}).Info("PollTweets started...")

	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		return nil, errors.New("consumer key/secret and access token/secret required")
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

	log.WithFields(log.Fields{
		"hashtag": query,
		"found":   len(search.Statuses),
	}).Info("PollTweets finished...")
	return &TwitterSearchResults{Query: query, SinceId: sinceID, Tweets: search.Statuses}, nil
}

// PublishTweets sends tweets via event publisher
func PublishTweets(ep services.EventPublisher, tweets []twitter.Tweet) error {
	log.Info("PublishTweets started...")
	events := make([]PollEvent, len(tweets))
	for i, tweet := range tweets {
		tm, err := time.Parse(TwitterTimeLayout, tweet.CreatedAt)
		if err != nil {
			return err
		}
		events[i] = PollEvent{
			EventEnvelope: services.EventEnvelope{
				Event:     services.PollTweetRunQuery,
				Timestamp: tm,
			},
			Subject: PollEventSubject{
				Buzzword:     tweet.Source,
				PartitionKey: tweet.Source,
			},
			Object: PollEventObject{
				TweetId:   tweet.ID,
				TweetText: tweet.Text,
			},
		}
	}
	ejsn := make([]services.EventJsoner, len(events))
	for i := range events {
		ejsn[i] = &events[i]
	}
	if err := ep.PublishEvents(ejsn); err != nil {
		return err
	}
	log.Info("PublishTweets finished...")
	return nil
}
