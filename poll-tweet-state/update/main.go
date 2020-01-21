package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	log "github.com/sirupsen/logrus"
	"os"
)

func Handler(state services.State) error {
	log.Infof("Got state to persist: %v", state)
	s := services.NewAwsDynamoDbStateStorer(os.Getenv("SERVERLESS_USER"), 1)
	return s.SetState(state)
}

func main() {
	lambda.Start(Handler)
}
