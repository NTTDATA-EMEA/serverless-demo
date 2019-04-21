package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event events.KinesisEvent) error {
	fmt.Printf("Received event: %+v\n", event)

	return nil
}

func main() {
	lambda.Start(handler)
}
