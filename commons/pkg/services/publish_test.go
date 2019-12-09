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

func TestEventPublisher(t *testing.T) {
	for _, eventPublisher := range getEventPublishers() {
		if err := eventPublisher.PublishEvent(EventEnvelope{
			Event:     POLL_TWEET_RUN_QUERY,
			Timestamp: time.Time{},
			Subject:   EventSubject{},
			Object:    EventObject{},
		}); err != nil {
			t.Errorf("Error publishing event: %s", err.Error())
		}
	}
}
