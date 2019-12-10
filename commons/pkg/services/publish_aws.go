package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"

	log "github.com/sirupsen/logrus"
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

func (ep *AwsEventPublisher) PublishEvent(event EventJsoner) error {
	return ep.PublishEvents([]EventJsoner{event})
}

func (ep *AwsEventPublisher) PublishEvents(events []EventJsoner) error {
	records := make([]*kinesis.PutRecordsRequestEntry, len(events))
	for i, event := range events {
		// jsn, err := json.Marshal(events[i])
		jsn, err := event.Marshal()
		if err != nil {
			return err
		}
		records[i] = &kinesis.PutRecordsRequestEntry{
			Data:         jsn,
			PartitionKey: aws.String(event.GetPartitionKey()),
		}
		log.Info(fmt.Sprintf("%s", jsn))
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
