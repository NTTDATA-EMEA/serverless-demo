package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"log"
	"os"
)

func handler(ctx context.Context, ke events.KinesisEvent) error {
	log.Printf("Received events, count %v\n", len(ke.Records))
	pes := make([]PollEvent, len(ke.Records))
	for i := range ke.Records {
		log.Printf("Kinesis data: %s", ke.Records[i].Kinesis.Data)
		if err := pes[i].Unmarshal(ke.Records[i].Kinesis.Data); err != nil {
			log.Printf("ERROR: Unmarshalling event: %s", err.Error())
			return err
		}
		log.Printf("Got PollEvent: %s, Buzzword: %s, Tweet: %s", pes[i].Event, pes[i].Subject.Buzzword, pes[i].Object.TweetText)
	}
	cbs := CollectBuzzwords(pes)
	if len(cbs) > 0 {
		streamName := os.Getenv("AWS_EVENT_STREAM_NAME")
		if streamName == "" {
			return errors.New("environment variable AWS_EVENT_STREAM_NAME is not defined")
		}
		ep := services.NewAwsEventPublisher(streamName)
		if err := PublishCollectBuzzwordAggregates(ep, cbs); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
