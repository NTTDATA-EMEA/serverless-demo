package services

import (
	"encoding/json"
	"fmt"
)

// LocalEventPublisher implements the EventPublisher interface
type LocalEventPublisher struct {
	StreamName string
}

// NewLocalEventPublisher returns a LocalEventPublisher implementation
func NewLocalEventPublisher(streamName string) EventPublisher {
	return &LocalEventPublisher{
		StreamName: streamName,
	}
}

// PublishEvent is a dummy implementation just logging the event
func (ep *LocalEventPublisher) PublishEvent(event *Event) error {
	json, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fmt.Printf("Serialised event: %s\n", json)

	return nil
}

// PublishEvents is a dummy implementation just logging the events
func (ep *LocalEventPublisher) PublishEvents(events []Event) error {
	for _, event := range events {
		ep.PublishEvent(&event)
	}
	return nil
}
