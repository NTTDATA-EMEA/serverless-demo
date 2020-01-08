package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"os"
)

func Handler() (services.State, error) {
	s := services.NewAwsDynamoDbStateStorer(os.Getenv("SERVERLESS_USER"), 1)
	return s.GetState()
}

func main() {
	lambda.Start(Handler)
}
