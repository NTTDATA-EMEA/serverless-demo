package main

import (
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
	// s := NewAwsStateStorer(o.Getenv("TWITTER_STATE_BUCKET"), "TwitterState.json")
	s := services.NewLocalStateStorer("/tmp", "TwitterState.json")
	_, err := PollAllTweets(s)
	if err != nil {
		t.Errorf(err.Error())
	}
}
