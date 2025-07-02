package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"github.com/isadoravieira/serverless-user-service/src/internal/domain/model"
)

func HandlerRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Método:", request.HTTPMethod)
	fmt.Println("Path:", request.Path)

	switch request.HTTPMethod {
	case "POST":
		if request.Path == "/user" {
			return createUser(request)
		}
	case "PUT":
		if strings.HasPrefix(request.Path, "/user/") {
			return updateUser(request)
		}
	case "GET":
		if request.Path == "/users" {
			return getUsers(request)
		}
		if strings.HasPrefix(request.Path, "/user/") {
			return getUserById(request)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 404,
		Body:       "Rota não encontrada",
	}, nil
}

func createUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Ta chegando na função createUser")

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

func updateUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := strings.TrimPrefix(request.Path, "/user/")

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: fmt.Sprintf("Usuário %s atualizado", id)}, nil
}

func getUsers(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: "Lista de usuários"}, nil
}

func getUserById(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := strings.TrimPrefix(request.Path, "/user/")

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: fmt.Sprintf("Usuário com ID %s", id)}, nil
}
