package services

// EventPublisher is technology agnostic interface to publish events
type EventPublisher interface {
	PublishEvent(event EventEnvelope) error
	PublishEvents(events []EventEnvelope) error
}
