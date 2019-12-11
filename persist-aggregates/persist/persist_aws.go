package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	log "github.com/sirupsen/logrus"
	"os"
)

type AwsDynamoDbAggregateStorer struct {
	TableName string
	Namespace string
	Version   int
}

func NewAwsDynamoDbAggregateStorer(namespace string, version int) *AwsDynamoDbAggregateStorer {
	return &AwsDynamoDbAggregateStorer{
		TableName: os.Getenv("BUZZWORD_AGGREGATES_TABLE"),
		Namespace: namespace,
		Version:   version,
	}
}

type AwsDynamoDbAggregateItem struct {
	BuzzwordNamespace string          `json:"buzzword_namespace"`
	Version           int             `json:"version"`
	Aggregate         *BuzzwordCounts `json:"aggregate"`
}

func (as AwsDynamoDbAggregateStorer) GetAggregate(buzzword string) (BuzzwordCounts, error) {
	panic("implement me")
}

func (as AwsDynamoDbAggregateStorer) GetAllAggregates() ([]BuzzwordCounts, error) {
	panic("implement me")
}

func (as AwsDynamoDbAggregateStorer) UpdateOrSetAggregate(ag *BuzzwordCounts) error {
	item := AwsDynamoDbAggregateItem{
		BuzzwordNamespace: ag.Keyword + as.Namespace,
		Version:           as.Version,
		Aggregate:         ag,
	}
	log.WithField("unmarshaled aggregate", item).Info("UpdateOrSetAggregate() to DynamoDB...")
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}
	log.WithField("marshaled aggregate", item).Info("UpdateOrSetAggregate() to DynamoDB...")
	svc, err := services.NewDynamoDbService()
	if err != nil {
		return err
	}
	out, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(as.TableName),
		Item:      av,
	})
	if err != nil {
		return err
	}
	log.WithField("written aggregate", out).Info("UpdateOrSetAggregate() to DynamoDB...")
	return nil
}
