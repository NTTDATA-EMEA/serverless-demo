package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/persist-aggregates/commons"
	"os"
)

func Handler() ([]commons.BuzzwordCounts, error) {
	s := commons.NewAwsDynamoDbAggregateStorer(os.Getenv("SERVERLESS_USER"), 1)
	return s.GetAllAggregates()
}

func main() {
	lambda.Start(Handler)
}
