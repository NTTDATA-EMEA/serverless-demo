package main

import (
	"os"
	"testing"
	"time"

	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestFindMaxSinceID(t *testing.T) {
	tweets, err := PollTweets("#cloud", 0)
	if err != nil {
		t.Errorf("Error polling tweets: %s", err.Error())
	}
	maxSinceID := int64(0)
	for _, tweet := range tweets.Tweets {
		t.Logf("Tweet ID: %d", tweet.ID)
		if tweet.ID > maxSinceID {
			maxSinceID = tweet.ID
		}
	}
	maxSinceIDUnderTest := findMaxSinceID(tweets.Tweets, maxSinceID)
	assert.Equal(t, maxSinceID, maxSinceIDUnderTest)
}

func TestPollTweets(t *testing.T) {
	tweets, err := PollTweets("#cloud", 0)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Tweets: %+v", tweets)
}

func TestPollTimeFormat(t *testing.T) {
	tm, err := time.Parse(TwitterTimeLayout, "Thu Apr 25 00:56:44 +0000 2019")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Logf("Parse result is %+v", tm)
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

func TestPublishTweets(t *testing.T) {
	tweets, err := PollTweets("#cloud", 0)
	if err != nil {
		t.Errorf(err.Error())
	}
	var ep services.EventPublisher
	if os.Getenv("AWS_INCLUDE_TESTS") == "1" {
		t.Logf("Using AWS publisher")
		ep = services.NewAwsEventPublisher(os.Getenv("AWS_EVENT_STREAM_NAME"))
	} else {
		t.Logf("Using local publisher")
		ep = services.NewLocalEventPublisher(os.Getenv("AWS_EVENT_STREAM_NAME"))
	}
	t.Logf("Publishing %d events", len(tweets.Tweets))
	if err := PublishTweets(ep, tweets.Tweets); err != nil {
		t.Errorf(err.Error())
	}
}
