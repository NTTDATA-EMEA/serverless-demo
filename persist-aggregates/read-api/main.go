package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/persist-aggregates/commons"
	"net/http"
)

type Response events.APIGatewayProxyResponse

func Handler(req events.APIGatewayProxyRequest) (Response, error) {
	keyword, ok := req.PathParameters["id"]
	if !ok {
		return Response{StatusCode: http.StatusBadRequest}, errors.New("keyword not defined in path parameters of the request")
	}
	result, err := commons.Invoke("read", keyword)
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
	lambda.Start(Handler)
}
