package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func NewDynamoDbService() (*dynamodb.DynamoDB, error) {
	sess, err := NewAwsSession()
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

func NewAwsSession() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
}
