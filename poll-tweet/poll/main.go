package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is generic return value of lambda
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// Handler receives and processes the request from the API gateway
func Handler(request events.CloudWatchEvent) (Response, error) {
	fmt.Printf("Received body: %s\n", request.ID)

	// Pseudo:
	// Get State
	// For reach: Get Tweets
	// Update SinceID
	// Set State

	return Response{
		Message: fmt.Sprintf("Processed request: %s", request.ID),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
