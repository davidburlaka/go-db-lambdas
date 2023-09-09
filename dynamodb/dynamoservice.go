package dynamodb

import (
	"fmt"
	"go-db-lambda/entitites"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func initDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	return dynamodb.New(sess)
}

func ReadFromDb(tableName string, username string) (entitites.ReadResponse, error) {
	var response entitites.ReadResponse

	svc := initDynamoSession()

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(username)},
		},
	}

	result, err := svc.GetItem(input)
	if err != nil {
		return entitites.ReadResponse{}, err
	}

	if result.Item == nil {
		return entitites.ReadResponse{}, fmt.Errorf("User does not exist")
	}

	if err := dynamodbattribute.UnmarshalMap(result.Item, &response); err != nil {
		return entitites.ReadResponse{}, err
	}

	return response, nil
}

func WriteToDb(tableName string, username string, data string) error {
	svc := initDynamoSession()

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {S: aws.String(username)},
			"data":     {S: aws.String(data)},
		},
	}

	_, err := svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
