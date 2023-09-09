package main

import (
	"go-db-lambda/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.ReadHandler)
}
