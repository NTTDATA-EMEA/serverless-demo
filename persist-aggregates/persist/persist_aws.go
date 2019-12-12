package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"os"
	"strconv"
	"time"
)

type TableKey map[string]*dynamodb.AttributeValue

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
	BuzzwordNamespace string         `json:"buzzword_namespace"`
	Version           int            `json:"version"`
	Aggregate         BuzzwordCounts `json:"aggregate"`
	UpdateCounter     int            `json:"update_counter"`
}

func (as AwsDynamoDbAggregateStorer) GetAggregate(buzzword string) (BuzzwordCounts, error) {
	empty := BuzzwordCounts{}
	svc, err := services.NewDynamoDbService()
	if err != nil {
		return empty, fmt.Errorf("GetAggregate.NewDynamoDbService()>%w", err)
	}
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		Key:       createTableKey(buzzword, as.Namespace, as.Version),
		TableName: aws.String(as.TableName),
	})
	if err != nil {
		return empty, fmt.Errorf("GetAggregate.GetItem()>%w", err)
	}
	item := AwsDynamoDbAggregateItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	fmt.Printf("Loaded from DB and unmarshalled: %+v", item)
	if err != nil {
		return empty, fmt.Errorf("GetAggregate.UnmarshalMap()>%w", err)
	}
	return item.Aggregate, nil
}

func (as AwsDynamoDbAggregateStorer) GetAllAggregates() ([]BuzzwordCounts, error) {
	panic("implement me")
}

func (as AwsDynamoDbAggregateStorer) UpdateOrSetAggregate(ag BuzzwordCounts) error {
	item := AwsDynamoDbAggregateItem{
		BuzzwordNamespace: ag.Keyword + as.Namespace,
		Version:           as.Version,
		Aggregate:         ag,
	}
	load, err := as.GetAggregate(ag.Keyword)
	if err != nil {
		return fmt.Errorf("UpdateOrSetAggregate.GetAggregate()>%w", err)
	}
	if load.Keyword != "" {
		err = item.AddBuzzwordCounts(load)
	}
	if err != nil {
		return fmt.Errorf("UpdateOrSetAggregate.AddBuzzwordCounts()>%w", err)
	}
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("UpdateOrSetAggregate.MarshalMap()>%w", err)
	}
	svc, err := services.NewDynamoDbService()
	if err != nil {
		return fmt.Errorf("UpdateOrSetAggregate.NewDynamoDbService()>%w", err)
	}
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(as.TableName),
		Item:      av,
	})
	if err != nil {
		return fmt.Errorf("UpdateOrSetAggregate:PutItem()>%w", err)
	}
	return nil
}

// AddBuzzwordCounts adds the counts from source to target
func (ai *AwsDynamoDbAggregateItem) AddBuzzwordCounts(source BuzzwordCounts) error {
	if ai.Aggregate.Keyword != source.Keyword {
		fmt.Printf("source: %+v; target: %+v", source, ai.Aggregate)
		return fmt.Errorf("keywords of aggregates are not the same")
	}
	for k := range source.Buzzwords {
		if _, ok := ai.Aggregate.Buzzwords[k]; !ok {
			ai.Aggregate.Buzzwords[k] = &BuzzwordCount{
				Keyword:    source.Keyword,
				Buzzword:   k,
				Count:      0,
				LastUpdate: time.Now(),
			}
		}
		ai.Aggregate.Buzzwords[k].Count += source.Buzzwords[k].Count
	}
	return nil
}

func createTableKey(buzzword string, namespace string, version int) TableKey {
	return TableKey{
		"buzzword_namespace": {
			S: aws.String(buzzword + namespace),
		},
		"version": {
			N: aws.String(strconv.Itoa(version)),
		},
	}
}
