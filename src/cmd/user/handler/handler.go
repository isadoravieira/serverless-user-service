package handler

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/isadoravieira/serverless-user-service/src/internal/domain/model"
)

func HandlerRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Ta chegando no handler")

	var user model.User

	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		fmt.Println("Deu erro pra fazer o unmarshal")
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
	}

	user.ID = uuid.New().String()

	// cria uma nova sessão na aws
	sess := session.Must(session.NewSession())

	// cria conexão com o dynamo
	svc := dynamodb.New(sess)

	// cria tabela e suas "colunas"
	input := &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(user.ID),
			},
			"name": {
				S: aws.String(user.Name),
			},
			"email": {
				S: aws.String(user.Email),
			},
			"password": {
				S: aws.String(user.Password),
			},
			"createdAt": {
				S: aws.String(user.CreatedAt),
			},
			"updatedAt": {
				S: aws.String(user.UpdatedAt),
			},
		},
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Deu erro pra criar item no dynamo")
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, err
	}

	body, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Deu erro pra fazer o marshal")
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}
