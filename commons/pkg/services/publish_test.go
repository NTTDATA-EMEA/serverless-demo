package services

import (
	"testing"
	"time"
)

func getEventPublishers() []EventPublisher {
	var eventPublishers []EventPublisher
	//eventPublishers =
	//	append(eventPublishers, NewAwsStateStorer(os.Getenv("TWITTER_STATE_BUCKET"), "testState.json"))
	eventPublishers =
		append(eventPublishers, NewLocalEventPublisher("test-stream"))
	return eventPublishers
}

/*
func getTestState() State {
	testState := make(State)
	testState["#cloud"] = 0
	testState["#ai"] = 0
	testState["#iot"] = 0
	return testState
}
*/

func getTestEvent() *Event {
	return &Event{
		ID:        "TEST-ID",
		Timestamp: time.Now(),
		Source:    "TEST",
		EventType: "TEST-EVENT",
		Payload:   getTestState(),
	}
}

func TestEventPublisher(t *testing.T) {
	for _, eventPublisher := range getEventPublishers() {
		if err := eventPublisher.PublishEvent(getTestEvent()); err != nil {
			t.Errorf("Error publishing event: %s", err.Error())
		}
	}
}
