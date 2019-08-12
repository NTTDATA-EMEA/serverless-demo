package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type AwsDynamoDbStateStorer struct {
	TableName string
	Namespace string
	Version   int
}

type AwsDynamoDbItem struct {
	Namespace string
	Version   int
	State     State
}

func (as *AwsDynamoDbStateStorer) GetState() (State, error) {
	svc, err := getDynamoDbService()
	if err != nil {
		return nil, err
	}
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(as.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Namespace": {
				S: aws.String(as.Namespace),
			},
			"Version": {
				N: aws.String(strconv.Itoa(as.Version)),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{
		"state read from db": result,
	}).Info("GetState() got from DynamoDB...")
	item := AwsDynamoDbItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{
		"state unmarshalling": item,
	}).Info("GetState() got from DynamoDB...")
	return item.State, nil
}

func (as *AwsDynamoDbStateStorer) SetState(state State) error {
	svc, err := getDynamoDbService()
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
	panic("implement me")
}

func NewAwsDynamoDbStateStorer(namespace string, version int) StateStorer {
	return &AwsDynamoDbStateStorer{
		TableName: os.Getenv("TWITTER_STATE_TABLE"),
		Namespace: namespace,
		Version:   version,
	}
}

func getDynamoDbService() (*dynamodb.DynamoDB, error) {
	sess, err := NewAwsSession()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}
