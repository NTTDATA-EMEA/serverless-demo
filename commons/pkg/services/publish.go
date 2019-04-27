package services

import "time"

// Event is the envelope of events used in the serverless-demo
type Event struct {
	ID        string
	Shard     string
	Timestamp time.Time
	Source    string
	EventType string
	Payload   interface{}
}

// EventPublisher is technology agnostic interface to publish events
type EventPublisher interface {
	PublishEvent(event *Event) error
	PublishEvents(events []Event) error
}
