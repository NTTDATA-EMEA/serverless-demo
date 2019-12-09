package services

import (
	"encoding/json"
	"errors"
	"time"
)

const (
	POLL_TWEET_RUN_QUERY         string = "POLL_TWEET_RUN_QUERY"
	COLLECT_BUZZWORDS_AGGREGATED string = "COLLECT_BUZZWORDS_AGGREGATED"
)

type EventJsoner interface {
	Marshal() ([]byte, error)
	Unmarshal(jsn []byte) error
	GetPartitionKey() (*string, error)
}

type EventEnvelope struct {
	Event     string       `json:"event"`
	Timestamp time.Time    `json:"timestamp"`
	Subject   EventSubject `json:"subject"`
	Object    EventObject  `json:"object"`
}

func (e EventEnvelope) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

func (e *EventEnvelope) Unmarshal(jsn []byte) error {
	return json.Unmarshal(jsn, e)
}

type EventSubject struct {
	Id         string            `json:"id"`
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}

type EventObject struct {
	Id         string                 `json:"id"`
	Name       string                 `json:"name"`
	Properties map[string]interface{} `json:"properties"`
}

func (e EventEnvelope) GetPartitionKey() (*string, error) {
	pk, ok := e.Subject.Properties["partitionKey"]
	if !ok {
		return nil, errors.New("partitionKey key not found in properties of the subject")
	}
	return &pk, nil
}
