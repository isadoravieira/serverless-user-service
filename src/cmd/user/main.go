package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/isadoravieira/serverless-user-service/src/cmd/user/handler"
)

func main() {
	lambda.Start(handler.HandlerRequest)
}
