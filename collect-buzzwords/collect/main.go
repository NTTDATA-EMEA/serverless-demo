package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
)

func handler(ctx context.Context, ke events.KinesisEvent) error {
	fmt.Printf("Received event: %+v\n", ke)
	for i := range ke.Records {
		var event services.Event
		if err := json.Unmarshal(ke.Records[i].Kinesis.Data, &event); err != nil {
			fmt.Printf("ERROR: Unmarshalling event: %s", err.Error())
			return err
		}
		fmt.Printf("Event: %+v from %v\n", event, ke.Records[i].Kinesis.ApproximateArrivalTimestamp)
		tweet, err := CreateTweetFromMap(event.Payload.(map[string]interface{}))
		if err != nil {
			fmt.Printf("ERROR: Creating tweet: %s", err.Error())
			return err
		}
		fmt.Printf("Tweet: %+v\n", tweet)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
