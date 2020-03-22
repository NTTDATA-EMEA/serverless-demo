package commons

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Response events.APIGatewayProxyResponse

func ResponseWithJson(body string) (*Response, error) {
	return &Response{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            body,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}
