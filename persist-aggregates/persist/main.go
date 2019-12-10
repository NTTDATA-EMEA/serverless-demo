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
	pes := make([]CollectEvent, len(ke.Records))
	for i := range ke.Records {
		fmt.Printf("Kinesis data: %s", ke.Records[i].Kinesis.Data)
		if err := pes[i].Unmarshal(ke.Records[i].Kinesis.Data); err != nil {
			fmt.Printf("ERROR: Unmarshalling event: %s", err.Error())
			return err
		}
		log.WithFields(log.Fields{
			"event":     pes[i].Event,
			"buzzword":  pes[i].Subject.Keyword,
			"aggregate": pes[i].Object.AggregatedBuzzwords,
		}).Info("Got CollectEvent")
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
