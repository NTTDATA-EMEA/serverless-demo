package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type TableKey map[string]*dynamodb.AttributeValue

type AwsDynamoDbStateStorer struct {
	TableName string
	Namespace string
	Version   int
}

type AwsDynamoDbItem struct {
	Namespace string `json:"namespace"`
	Version   int    `json:"version"`
	State     State  `json:"state"`
}

func (as *AwsDynamoDbStateStorer) GetState() (State, error) {
	svc, err := NewDynamoDbService()
	if err != nil {
		return nil, err
	}
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(as.TableName),
		Key:       createTableKey(as.Namespace, as.Version),
	})
	if err != nil {
		return nil, err
	}
	item := AwsDynamoDbItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}
	log.WithField("state unmarshalling", item).Info("GetState() got from DynamoDB...")
	return item.State, nil
}

func (as *AwsDynamoDbStateStorer) SetState(state State) error {
	svc, err := NewDynamoDbService()
	if err != nil {
		return err
	}
	item := AwsDynamoDbItem{
		Namespace: as.Namespace,
		Version:   as.Version,
		State:     state,
	}
	log.WithFields(log.Fields{
		"unmarshaled state": item,
	}).Info("SetState() to DynamoDB...")
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"marshaled state": av,
	}).Info("SetState() to DynamoDB...")
	out, err := svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(as.TableName),
		Item:      av,
	})
	if err != nil {
		return err
	}
	log.WithFields(log.Fields{
		"written state": out,
	}).Info("SetState() to DynamoDB...")
	return nil
}

func (as *AwsDynamoDbStateStorer) DeleteState() error {
	svc, err := NewDynamoDbService()
	if err != nil {
		return err
	}
	_, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(as.TableName),
		Key:       createTableKey(as.Namespace, as.Version),
	})
	if err != nil {
		return err
	}
	return nil
}

func NewAwsDynamoDbStateStorer(namespace string, version int) StateStorer {
	return &AwsDynamoDbStateStorer{
		TableName: os.Getenv("TWITTER_STATE_TABLE"),
		Namespace: namespace,
		Version:   version,
	}
}

func createTableKey(namespace string, version int) TableKey {
	return TableKey{
		"namespace": {
			S: aws.String(namespace),
		},
		"version": {
			N: aws.String(strconv.Itoa(version)),
		},
	}
}
