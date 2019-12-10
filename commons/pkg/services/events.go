package services

import (
	"time"
)

const (
	PollTweetRunQuery          string = "POLL_TWEET_RUN_QUERY"
	CollectBuzzwordsAggregated string = "COLLECT_BUZZWORDS_AGGREGATED"
)

type EventJsoner interface {
	Marshal() ([]byte, error)
	Unmarshal(jsn []byte) error
	GetPartitionKey() string
}

type EventEnvelope struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
