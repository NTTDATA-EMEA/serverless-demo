package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/persist-aggregates/commons"
	"os"
)

func Handler(keyword string) (commons.BuzzwordCounts, error) {
	s := commons.NewAwsDynamoDbAggregateStorer(os.Getenv("SERVERLESS_USER"), 1)
	qkeyword := "#" + keyword
	r, err := s.GetAggregate(qkeyword)
	if err != nil {
		return commons.BuzzwordCounts{}, err
	}
	if r.Keyword == "" {
		r.Keyword = qkeyword
	}
	return r, nil
}

func main() {
	lambda.Start(Handler)
}
