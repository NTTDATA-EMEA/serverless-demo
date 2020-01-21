package main

import (
	"encoding/json"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
)

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

func (pe PollEvent) GetPartitionKey() string {
	return pe.Subject.PartitionKey
}

func (pe PollEvent) Marshal() ([]byte, error) {
	return json.Marshal(pe)
}

func (pe *PollEvent) Unmarshal(jsn []byte) error {
	return json.Unmarshal(jsn, pe)
}
