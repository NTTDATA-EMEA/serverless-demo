package services

import (
	"encoding/json"
	"time"
)

const (
	POLL_TWEET_RUN_QUERY string = "POLL_TWEET_RUN_QUERY"
)

type EventJsoner interface {
	Marshal() ([]byte, error)
	Unmarshal(jsn []byte) error
	GetBuzzword() string
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
	Id         string            `json:"id"`
	Name       string            `json:"name"`
	Properties map[string]string `json:"properties"`
}
