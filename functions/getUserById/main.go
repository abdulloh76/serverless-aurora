package main

import (
	"os"

	"github.com/abdulloh76/serverless-aurora/domain"
	"github.com/abdulloh76/serverless-aurora/handlers"
	"github.com/abdulloh76/serverless-aurora/store"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	postgresDSN, ok := os.LookupEnv("POSTGRES_URI")
	if !ok {
		panic("Need POSTGRES_URI environment variable")
	}

	postgreDB := store.NewPostgresDBStore(postgresDSN)
	domain := domain.NewUsersDomain(postgreDB)
	handler := handlers.NewAPIGatewayV2Handler(domain)
	lambda.Start(handler.GetHandler)
}
