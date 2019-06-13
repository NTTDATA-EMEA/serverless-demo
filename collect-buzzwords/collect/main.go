package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"

	log "github.com/sirupsen/logrus"
)

func handler(ctx context.Context, ke events.KinesisEvent) error {
	log.WithFields(log.Fields{
		"count": len(ke.Records),
	}).Info("Received events")
	for i := range ke.Records {
		var event services.Event
		if err := json.Unmarshal(ke.Records[i].Kinesis.Data, &event); err != nil {
			fmt.Printf("ERROR: Unmarshalling event: %s", err.Error())
			return err
		}
		log.WithFields(log.Fields{
			"id": event.ID,
			"type": event.EventType,
			"hashtag": event.Source,
		}).Info("Event header")
		// fmt.Printf("Event: %+v from %v\n", event, ke.Records[i].Kinesis.ApproximateArrivalTimestamp)
		tweet, err := CreateTweetFromMap(event.Payload.(map[string]interface{}))
		if err != nil {
			fmt.Printf("ERROR: Creating tweet: %s", err.Error())
			return err
		}
		log.WithFields(log.Fields{
			"body": tweet,
		}).Info("Tweet")
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
