package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"os"

	log "github.com/sirupsen/logrus"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler() (Response, error) {
	var buf bytes.Buffer
	// s := services.NewAwsStateStorer(os.Getenv("TWITTER_STATE_BUCKET"), os.Getenv("TWITTER_STATE_FILE"))
	s := services.NewAwsDynamoDbStateStorer(os.Getenv("SERVERLESS_USER"), 1)
	state, err := s.GetState()
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	sjson, err := json.Marshal(state)
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, sjson)
	log.Infof("State read: %v", buf.String())

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
