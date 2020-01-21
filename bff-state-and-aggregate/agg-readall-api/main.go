package main

import (
	"fmt"
	lmb "github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/bff-state-and-aggregate/commons"

	"net/http"
)

func Handler() (commons.Response, error) {
	result, err := commons.Invoke(commons.PERSIST_AGGREGATES_MODULE, "readall", nil)
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
