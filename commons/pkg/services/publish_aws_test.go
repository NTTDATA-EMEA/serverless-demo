package services

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func TestKinesisConnection(t *testing.T) {
	streamName := os.Getenv("AWS_EVENT_STREAM_NAME")
	if streamName == "" {
		t.Errorf("Stream name not defined, try export AWS_EVENT_STREAM=arn...")
	}
	svc := kinesis.New(session.New())
	_, err := svc.PutRecords(&kinesis.PutRecordsInput{
		Records: []*kinesis.PutRecordsRequestEntry{
			{
				Data:         []byte("x"),
				PartitionKey: aws.String("b"),
			},
		},
		StreamName: aws.String(streamName),
	})
	if err != nil {
		t.Errorf("Error putting record: %s", err.Error())
	}

}
