package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	lmb "github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/bff-state-and-aggregate/commons"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Handler(req events.APIGatewayProxyRequest) (commons.Response, error) {
	log.Infof("Got request to persist: %v", req.Body)
	var state services.State
	if err := json.Unmarshal([]byte(req.Body), &state); err != nil {
		return commons.Response{StatusCode: 404}, fmt.Errorf("handler.json.unmarshal error: %w", err)
	}
	result, err := commons.Invoke(commons.POLL_TWEET_STATE_MODULE, "update", state)
	if err != nil {
		return commons.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("handler.commons.invoke error: %w", err)
	}
	if result.Payload != nil {
		return commons.Response{StatusCode: http.StatusInternalServerError}, errors.New(string(result.Payload))
	}

	return commons.ResponseWithJson("State updates successfully!"), nil
}

func main() {
	lmb.Start(Handler)
}
