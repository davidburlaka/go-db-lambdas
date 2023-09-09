package handlers

import (
	"context"
	"go-db-lambda/dynamodb"
	"go-db-lambda/entitites"
)

func WriteHandler(ctx context.Context, event entitites.WriteEvent) (string, error) {
	tableName := entitites.TableNameBase + event.EntityId

	err := dynamodb.WriteToDb(tableName, event.Username, event.Data)
	if err != nil {
		return "", err
	}

	return "Data has been saved", nil
}
