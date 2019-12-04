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
func (ep *LocalEventPublisher) PublishEvent(event EventJsoner) error {
	return ep.PublishEvents([]EventJsoner{event})
}

// PublishEvents is a dummy implementation just logging the events (highly inefficient ;-)
func (ep *LocalEventPublisher) PublishEvents(events []EventJsoner) error {
	for _, event := range events {
		jsn, err := json.Marshal(event)
		if err != nil {
			return err
		}
		fmt.Printf("Serialised event: %s\n", jsn)
	}
	return nil
}
