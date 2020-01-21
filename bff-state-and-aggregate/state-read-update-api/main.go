package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	lmb "github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/bff-state-and-aggregate/commons"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	log "github.com/sirupsen/logrus"

	"net/http"
)

func Handler(req events.APIGatewayProxyRequest) (commons.Response, error) {
	var res *commons.Response
	switch req.HTTPMethod {
	case "GET":
		log.Infof("Got request to read state...")
		result, err := commons.Invoke(commons.POLL_TWEET_STATE_MODULE, "read", nil)
		if err != nil {
			return commons.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("handler.commons.invoke error: %w", err)
		}
		res, err = commons.ResponseWithJson(string(result.Payload))
		if err != nil {
			return commons.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("handler.commons.invoke error: %w", err)
		}
	case "PUT":
		log.Infof("Got request to persist state: %v", req.Body)
		var state services.State
		if err := json.Unmarshal([]byte(req.Body), &state); err != nil {
			return commons.Response{StatusCode: 404}, fmt.Errorf("handler.json.unmarshal error: %w", err)
		}
		result, err := commons.Invoke(commons.POLL_TWEET_STATE_MODULE, "update", state)
		if err != nil {
			return commons.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("handler.commons.invoke error: %w", err)
		}
		if string(result.Payload) != "null" {
			return commons.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("handler.commons.invoke error: %v", result.Payload)
		}
		res, err = commons.ResponseWithJson("State updated successfully!")
		if err != nil {
			return commons.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("handler.commons.invoke error: %w", err)
		}
	}
	return *res, nil
}

func main() {
	lmb.Start(Handler)
}
