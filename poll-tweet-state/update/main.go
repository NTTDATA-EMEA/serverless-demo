package main

import (
	"bytes"
	"encoding/json"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req events.APIGatewayProxyRequest) (Response, error) {
	var buf bytes.Buffer

	var state services.State
	if err := json.Unmarshal([]byte(req.Body), &state); err != nil {
		return Response{StatusCode: 404}, err
	}

	// s := services.NewAwsStateStorer(os.Getenv("TWITTER_STATE_BUCKET"), os.Getenv("TWITTER_STATE_FILE"))
	s := services.NewAwsDynamoDbStateStorer(os.Getenv("SERVERLESS_USER"), 1)
	if err := s.SetState(state); err != nil {
		return Response{StatusCode: 404}, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"message": "Upload executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	log.Infof("State uploaded: %s", req.Body)

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
