package services

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// AwsEventPublisher implements the EventPublisher interface
type AwsEventPublisher struct {
	StreamName string
}

// NewAwsEventPublisher returns an AwsEventPublisher implementation
func NewAwsEventPublisher(streamName string) EventPublisher {
	return &AwsEventPublisher{
		StreamName: streamName,
	}
}

// PublishEvent is a dummy implementation just logging the event
func (ep *AwsEventPublisher) PublishEvent(event Event) error {
	ep.PublishEvents([]Event{event})
	return nil
}

// PublishEvents is a dummy implementation just logging the events
func (ep *AwsEventPublisher) PublishEvents(events []Event) error {
	var records []*kinesis.PutRecordsRequestEntry
	for i := range events {
		json, err := json.Marshal(events[i])
		if err != nil {
			return err
		}
		records = append(records, &kinesis.PutRecordsRequestEntry{
			Data:         json,
			PartitionKey: aws.String(events[i].Shard),
		})
	}
	svc := kinesis.New(session.New())
	if _, err := svc.PutRecords(&kinesis.PutRecordsInput{
		Records:    records,
		StreamName: aws.String(ep.StreamName),
	}); err != nil {
		return err
	}
	return nil
}
