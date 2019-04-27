package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
)

// Response is generic return value of lambda
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// Handler receives and processes the request from the API gateway
func Handler(request events.CloudWatchEvent) (Response, error) {
	fmt.Printf("Received body: %s\n", request.ID)

	s := services.NewAwsStateStorer(os.Getenv("TWITTER_STATE_BUCKET"), os.Getenv("TWITTER_STATE_FILE"))
	_, err := PollAllTweets(s)
	if err != nil {
		return Response{
			Message: fmt.Sprintf("Processed request: %s with error %s", request.ID, err.Error()),
			Ok:      false,
		}, err
	}
	// PublishAllTweets(tweets)

	return Response{
		Message: fmt.Sprintf("Processed request: %s", request.ID),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
