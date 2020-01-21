package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	lmb "github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/bff-state-and-aggregate/commons"
	"net/http"
)

func Handler(req events.APIGatewayProxyRequest) (commons.Response, error) {
	keyword, ok := req.PathParameters["id"]
	if !ok {
		return commons.Response{StatusCode: http.StatusBadRequest}, errors.New("handler error: keyword not defined in path parameters of the request")
	}
	result, err := commons.Invoke(commons.PERSIST_AGGREGATES_MODULE, "read", keyword)
	if err != nil {
		return commons.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("handler.commons.invoke error: %w", err)
	}
	res, err := commons.ResponseWithJson(string(result.Payload))
	if err != nil {
		return commons.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("handler.commons.invoke error: %w", err)
	}
	return *res, nil
}

func main() {
	lmb.Start(Handler)
}
