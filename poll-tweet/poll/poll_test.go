package main

import (
	"os"
	"testing"

	"github.com/okoeth/serverless-demo/commons/pkg/services"
)

func TestFindMaxSinceID(t *testing.T) {
	tweets, err := PollTweets("#cloud", 0)
	if err != nil {
		t.Errorf("Error polling tweets: %s", err.Error())
	}
	maxSinceID := int64(0)
	for _, tweet := range tweets {
		t.Logf("Tweet ID: %d", tweet.ID)
		if tweet.ID > maxSinceID {
			maxSinceID = tweet.ID
		}
	}
	if maxSinceID != findMaxSinceID(tweets, 0) {
		t.Errorf("Difference in maxSinceID")
	}
	t.Logf("MaxSinceID is %d", maxSinceID)
}

func TestPollTweets(t *testing.T) {
	_, err := PollTweets("#cloud", 0)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPollAllTweets(t *testing.T) {
	var s services.StateStorer
	if os.Getenv("AWS_INCLUDE_TESTS") == "1" {
		s = services.NewAwsStateStorer(os.Getenv("TWITTER_STATE_BUCKET"), os.Getenv("TWITTER_STATE_FILE"))
	} else {
		s = services.NewLocalStateStorer("/tmp", os.Getenv("TWITTER_STATE_FILE"))
	}
	_, err := PollAllTweets(s)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPollPublishTweets(t *testing.T) {
}

func TestPollPublishAllTweets(t *testing.T) {
}
