package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/persist-aggregates/commons"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type Response events.APIGatewayProxyResponse

func Handler(req events.APIGatewayProxyRequest) (Response, error) {
	keyword, ok := req.PathParameters["id"]
	if !ok {
		return Response{StatusCode: http.StatusBadRequest}, errors.New("keyword not defined in path parameters of the request")
	}
	s := commons.NewAwsDynamoDbAggregateStorer(os.Getenv("SERVERLESS_USER"), 1)
	agg, err := s.GetAggregate("#" + keyword)
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}
	var buf bytes.Buffer
	sjson, err := json.Marshal(agg)
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}
	json.HTMLEscape(&buf, sjson)
	log.Infof("Aggregate read: %v", buf.String())

	resp := Response{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
