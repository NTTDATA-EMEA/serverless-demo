package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/okoeth/serverless-demo/persist-aggregates/commons"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type Response events.APIGatewayProxyResponse

func Handler() (Response, error) {
	s := commons.NewAwsDynamoDbAggregateStorer(os.Getenv("SERVERLESS_USER"), 1)
	agg, err := s.GetAllAggregates()
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}
	var buf bytes.Buffer
	sjson, err := json.Marshal(agg)
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError}, err
	}
	json.HTMLEscape(&buf, sjson)
	log.Infof("Aggregates read: %v", buf.String())

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
