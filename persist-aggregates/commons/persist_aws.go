package commons

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/okoeth/serverless-demo/commons/pkg/services"
	"log"
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

// GetAggregate retrieves one aggregate by buzzword as key
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
	if err != nil {
		return empty, fmt.Errorf("GetAggregate.UnmarshalMap()>%w", err)
	}
	return item.Aggregate, nil
}

// GetAllAggregates retrieves all aggregates from DB
func (as AwsDynamoDbAggregateStorer) GetAllAggregates() ([]BuzzwordCounts, error) {
	svc, err := services.NewDynamoDbService()
	if err != nil {
		return nil, fmt.Errorf("GetAllAggregates.NewDynamoDbService()>%w", err)
	}
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(as.TableName),
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllAggregates.Scan()>%w", err)
	}
	bcs := make([]BuzzwordCounts, len(result.Items))
	for i := range result.Items {
		item := AwsDynamoDbAggregateItem{}
		err = dynamodbattribute.UnmarshalMap(result.Items[i], &item)
		if err != nil {
			return nil, fmt.Errorf("GetAllAggregates.UnmarshalMap()>%w", err)
		}
		bcs[i] = item.Aggregate
	}
	return bcs, nil
}

// UpdateOrSetAggregate updates the aggregate with new counts or buzzwords
// or creates new aggregate
func (as AwsDynamoDbAggregateStorer) UpdateOrSetAggregate(ag BuzzwordCounts) error {
	log.Printf("UpdateOrSetAggregate: keyword: %s, namespace: %s", ag.Keyword, as.Namespace)
	item := AwsDynamoDbAggregateItem{
		BuzzwordNamespace: ag.Keyword + as.Namespace,
		Version:           as.Version,
		Aggregate:         ag,
	}
	load, err := as.GetAggregate(ag.Keyword)
	log.Printf("GetAggregate: load: %v", load)
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
	log.Printf("Item to put in table %s: %v", as.TableName, item)
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
