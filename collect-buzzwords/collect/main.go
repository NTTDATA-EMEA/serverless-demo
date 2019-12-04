package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

func handler(ctx context.Context, ke events.KinesisEvent) error {
	log.WithFields(log.Fields{
		"count": len(ke.Records),
	}).Info("Received events")
	for i := range ke.Records {
		var event PollEvent
		if err := event.Unmarshal(ke.Records[i].Kinesis.Data); err != nil {
			fmt.Printf("ERROR: Unmarshalling event: %s", err.Error())
			return err
		}
		log.WithFields(log.Fields{
			"id":    event.Object.Id,
			"name":  event.Object.Name,
			"shard": event.Subject.Properties["shard"],
			"tweet": event.Object.Properties["body"],
		}).Info("Got PollEvent")

	}
	return nil
}

func main() {
	lambda.Start(handler)
}
