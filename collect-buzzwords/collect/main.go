package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	log "github.com/sirupsen/logrus"
	"os"
)

func handler(ctx context.Context, ke events.KinesisEvent) error {
	log.WithFields(log.Fields{
		"count": len(ke.Records),
	}).Info("Received events")
	pes := make([]PollEvent, len(ke.Records))
	for i := range ke.Records {
		if err := pes[i].Unmarshal(ke.Records[i].Kinesis.Data); err != nil {
			fmt.Printf("ERROR: Unmarshalling event: %s", err.Error())
			return err
		}
		log.WithFields(log.Fields{
			"id":       pes[i].Object.Id,
			"name":     pes[i].Object.Name,
			"buzzword": pes[i].GetBuzzword(),
			"tweet":    pes[i].GetTweetText(),
		}).Info("Got PollEvent")
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
