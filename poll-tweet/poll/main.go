package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	log "github.com/sirupsen/logrus"
)

// Response is generic return value of lambda
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func makeError(requestID string, err error) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Processed request: %s with error %s", requestID, err.Error()),
		Ok:      false,
	}, err
}

func Handler(request events.CloudWatchEvent) (Response, error) {
	log.WithFields(log.Fields{"ID": request.ID}).Info("Event received")
	var s services.StateStorer

	/* Variant storing the state in S3
	stateBucket := os.Getenv("TWITTER_STATE_BUCKET")
	if stateBucket == "" {
		return makeError(request.ID, errors.New(
			"environment variable TWITTER_STATE_BUCKET is not defined"))

	}
	stateFile := os.Getenv("TWITTER_STATE_FILE")
	if stateFile == "" {
		return makeError(request.ID, errors.New(
			"environment variable TWITTER_STATE_FILE is not defined"))
	}
	s = services.NewAwsStateStorer(stateBucket, stateFile)
	*/

	s = services.NewAwsDynamoDbStateStorer(os.Getenv("SERVERLESS_USER"), 1)
	tweets, err := PollAllTweets(s)
	if err != nil {
		return makeError(request.ID, err)
	}
	log.WithFields(log.Fields{
		"count": len(tweets),
	}).Info("Publishing new tweets")
	if len(tweets) > 0 {
		streamName := os.Getenv("AWS_EVENT_STREAM_NAME")
		if streamName == "" {
			return makeError(request.ID, errors.New(
				"environment variable AWS_EVENT_STREAM_NAME is not defined"))
		}
		ep := services.NewAwsEventPublisher(streamName)
		if err := PublishTweets(ep, tweets); err != nil {
			return makeError(request.ID, err)
		}
	}
	return Response{
		Message: fmt.Sprintf("Processed request: %s", request.ID),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
