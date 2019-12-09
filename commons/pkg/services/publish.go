package services

// EventPublisher is technology agnostic interface to publish events
type EventPublisher interface {
	PublishEvent(event EventJsoner) error
	PublishEvents(events []EventJsoner) error
}
