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

func (ep *AwsEventPublisher) PublishEvent(event EventEnvelope) error {
	return ep.PublishEvents([]EventEnvelope{event})
}

func (ep *AwsEventPublisher) PublishEvents(events []EventEnvelope) error {
	records := make([]*kinesis.PutRecordsRequestEntry, len(events))
	for i := range events {
		jsn, err := json.Marshal(events[i])
		if err != nil {
			return err
		}
		records[i] = &kinesis.PutRecordsRequestEntry{
			Data:         jsn,
			PartitionKey: aws.String(events[i].Subject.Properties["shard"]),
		}
	}
	s, err := session.NewSession()
	if err != nil {
		return err
	}
	svc := kinesis.New(s)
	if _, err := svc.PutRecords(&kinesis.PutRecordsInput{
		Records:    records,
		StreamName: aws.String(ep.StreamName),
	}); err != nil {
		return err
	}
	return nil
}
