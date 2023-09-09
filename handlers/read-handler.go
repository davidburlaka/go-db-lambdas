package handlers

import (
	"context"
	"go-db-lambda/dynamodb"
	"go-db-lambda/entitites"
)

func ReadHandler(ctx context.Context, event entitites.ReadEvent) (*entitites.ReadResponse, error) {
	tableName := entitites.TableNameBase + event.EntityId

	response, err := dynamodb.ReadFromDb(tableName, event.Username)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
