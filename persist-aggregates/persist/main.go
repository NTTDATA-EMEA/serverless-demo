package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/persist-aggregates/commons"
	"os"
)

func handler(ctx context.Context, ke events.KinesisEvent) error {
	storer := commons.NewAwsDynamoDbAggregateStorer(os.Getenv("SERVERLESS_USER"), 1)
	pes := make([]commons.CollectEvent, len(ke.Records))
	for i := range ke.Records {
		fmt.Printf("Kinesis data: %s", ke.Records[i].Kinesis.Data)
		if err := pes[i].Unmarshal(ke.Records[i].Kinesis.Data); err != nil {
			return fmt.Errorf("ERROR in handler.UpdateOrSetAggregate()>%w", err)
		}
		if err := storer.UpdateOrSetAggregate(pes[i].Object.AggregatedBuzzwords); err != nil {
			return fmt.Errorf("ERROR in handler.UpdateOrSetAggregate()>%w", err)
		}
	}
	return nil
}

func main() {
	lambda.Start(handler)
}
