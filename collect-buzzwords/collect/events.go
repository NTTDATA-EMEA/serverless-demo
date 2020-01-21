package main

import (
	"encoding/json"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
)

type CollectEvent struct {
	services.EventEnvelope
	Subject CollectEventSubject `json:"subject"`
	Object  CollectEventObject  `json:"object"`
}

type CollectEventSubject struct {
	PartitionKey string `json:"partition_key"`
	Keyword      string `json:"keyword"`
}

type CollectEventObject struct {
	AggregatedBuzzwords *BuzzwordCounts `json:"aggregated_buzzwords"`
}

func (ce CollectEvent) GetPartitionKey() string {
	return ce.Subject.PartitionKey
}

func (ce CollectEvent) Marshal() ([]byte, error) {
	return json.Marshal(ce)
}

func (ce *CollectEvent) Unmarshal(jsn []byte) error {
	return json.Unmarshal(jsn, ce)
}

type PollEvent struct {
	services.EventEnvelope
	Subject PollEventSubject `json:"subject"`
	Object  PollEventObject  `json:"object"`
}

type PollEventSubject struct {
	Buzzword     string `json:"buzzword"`
	PartitionKey string `json:"partition_key"`
}

type PollEventObject struct {
	TweetId   int64  `json:"tweet_id"`
	TweetText string `json:"tweet_text"`
}

func (pe *PollEvent) Unmarshal(jsn []byte) error {
	return json.Unmarshal(jsn, pe)
}
