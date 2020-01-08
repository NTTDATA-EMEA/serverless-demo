package main

import (
	"github.com/aws/aws-lambda-go/events"
	lmb "github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/persist-aggregates/commons"
	"net/http"
)

type Response events.APIGatewayProxyResponse

func Handler() (Response, error) {
	result, err := commons.Invoke("readall", nil)
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}
	resp := Response{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            string(result.Payload),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	return resp, nil
}

func main() {
	lmb.Start(Handler)
}
