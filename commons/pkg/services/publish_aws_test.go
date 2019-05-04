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
		t.Errorf("Stream name not defined, try export AWS_EVENT_STREAM=...")
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

func TestDumpKinesis(t *testing.T) {
	streamName := os.Getenv("AWS_EVENT_STREAM_NAME")
	if streamName == "" {
		t.Errorf("Stream name not defined, try export AWS_EVENT_STREAM_NAME=...")
	}
	svc := kinesis.New(session.New())
	ds, err := svc.DescribeStream(&kinesis.DescribeStreamInput{
		StreamName: aws.String(streamName),
	})
	if err != nil {
		t.Errorf("Error putting record: %s", err.Error())
	}
	t.Logf("Describe Stream: %+v", ds)
	for _, shard := range ds.StreamDescription.Shards {
		si, err := svc.GetShardIterator(&kinesis.GetShardIteratorInput{
			ShardId:           shard.ShardId,
			ShardIteratorType: aws.String("TRIM_HORIZON"),
			StreamName:        aws.String(streamName),
		})
		if err != nil {
			t.Errorf("Error putting record: %s", err.Error())
		}
		t.Logf("Shard Iterator: %+v", si)
		r, err := svc.GetRecords(&kinesis.GetRecordsInput{
			ShardIterator: si.ShardIterator,
		})
		if err != nil {
			t.Errorf("Error putting record: %s", err.Error())
		}
		t.Logf("Records: %+v", r)
	}
}
