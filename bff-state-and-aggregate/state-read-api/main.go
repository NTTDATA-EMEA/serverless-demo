package main

import (
	lmb "github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/bff-state-and-aggregate/commons"

	"net/http"
)

func Handler() (commons.Response, error) {
	result, err := commons.Invoke(commons.POLL_TWEET_STATE_MODULE, "read", nil)
	if err != nil {
		return commons.Response{StatusCode: http.StatusInternalServerError}, err
	}

	return commons.ResponseWithJson(string(result.Payload)), nil
}

func main() {
	lmb.Start(Handler)
}
