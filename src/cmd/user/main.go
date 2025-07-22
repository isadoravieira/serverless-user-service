package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/isadoravieira/serverless-user-service/src/cmd/user/handler"
)

func main() {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	h := handler.NewUserHandler(svc)

	lambda.Start(h.Handle)
}
