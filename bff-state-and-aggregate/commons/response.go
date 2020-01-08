package commons

import (
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Response events.APIGatewayProxyResponse

func ResponseWithJson(body string) Response {
	return Response{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
