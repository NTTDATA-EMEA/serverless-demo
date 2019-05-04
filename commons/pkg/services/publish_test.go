package services

import (
	"os"
	"testing"
	"time"
)

func getEventPublishers() []EventPublisher {
	var eventPublishers []EventPublisher
	if os.Getenv("AWS_INCLUDE_TESTS") == "1" {
		eventPublishers =
			append(eventPublishers, NewAwsEventPublisher(os.Getenv("AWS_EVENT_STREAM_NAME")))
	}
	eventPublishers =
		append(eventPublishers, NewLocalEventPublisher(os.Getenv("AWS_EVENT_STREAM_NAME")))
	return eventPublishers
}

func getTestEvent() Event {
	return Event{
		ID:        "TEST-ID",
		Shard:     "SHARD1",
		Timestamp: time.Now(),
		Source:    "TEST",
		EventType: "TEST-EVENT",
		Payload:   []string{"Hello", "World"},
	}
}

func TestEventPublisher(t *testing.T) {
	for _, eventPublisher := range getEventPublishers() {
		if err := eventPublisher.PublishEvent(getTestEvent()); err != nil {
			t.Errorf("Error publishing event: %s", err.Error())
		}
	}
}
