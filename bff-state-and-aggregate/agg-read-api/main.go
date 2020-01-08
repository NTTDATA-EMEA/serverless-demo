package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	lmb "github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/bff-state-and-aggregate/commons"
	"net/http"
)

func Handler(req events.APIGatewayProxyRequest) (commons.Response, error) {
	keyword, ok := req.PathParameters["id"]
	if !ok {
		return commons.Response{StatusCode: http.StatusBadRequest}, errors.New("keyword not defined in path parameters of the request")
	}
	result, err := commons.Invoke(commons.PERSIST_AGGREGATES_MODULE, "read", keyword)
	if err != nil {
		return commons.Response{StatusCode: http.StatusInternalServerError}, err
	}

	return commons.ResponseWithJson(string(result.Payload)), nil
}

func main() {
	lmb.Start(Handler)
}
