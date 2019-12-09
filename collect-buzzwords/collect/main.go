package main

import (
	"context"
	"fmt"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
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
	for i := range cbs {
		event := CollectEvent{
			EventEnvelope: services.EventEnvelope{
				Event:     services.COLLECT_BUZZWORDS_AGGREGATED,
				Timestamp: time.Time{},
				Subject: services.EventSubject{
					Id:   i,
					Name: cbs[i].Keyword,
				},
				Object: services.EventObject{
					Id:   i,
					Name: "aggregate",
					Properties: map[string]interface{}{
						"buzzwords": cbs[i].Buzzwords,
					},
				},
			},
		}
		jsn, err := event.Marshal()
		if err != nil {
			return err
		}
		log.WithField("buzzword", fmt.Sprintf("%s", jsn)).Info("Marshalled Collect Event...")
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
