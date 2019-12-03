package services

import "time"

const (
	POLL_TWEET_RUN_QUERY string = "POLL_TWEET_RUN_QUERY"
)

type EventEnvelope struct {
	Event     string       `json:"event"`
	Timestamp time.Time    `json:"timestamp"`
	Subject   EventSubject `json:"subject"`
	Object    EventObject  `json:"object"`
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
