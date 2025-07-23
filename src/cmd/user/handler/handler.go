package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/isadoravieira/serverless-user-service/src/internal/application/service"
	"github.com/isadoravieira/serverless-user-service/src/internal/domain/model"
	"github.com/isadoravieira/serverless-user-service/src/pkg/responses"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(svc *dynamodb.DynamoDB) *UserHandler {
	return &UserHandler{
		UserService: service.NewUserService(svc),
	}
}

func (h *UserHandler) Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Método:", request.HTTPMethod)
	fmt.Println("Path:", request.Path)

	switch request.HTTPMethod {
	case "POST":
		if request.Path == "/user" {
			return h.createUser(request)
		}
	case "PUT":
		if strings.HasPrefix(request.Path, "/user/") {
			return h.updateUser(request)
		}
	case "GET":
		if request.Path == "/users" {
			return h.getUsers(request)
		}
		if strings.HasPrefix(request.Path, "/user/") {
			return h.getUserById(request)
		}
	}

	return responses.DomainJSON(404, "Route not found"), nil
}

func (h *UserHandler) createUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user model.User

	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		return responses.DomainError(500, err), err
	}

	if err = user.PrepareUser("register"); err != nil {
		return responses.DomainError(500, err), err
	}

	createdUser, err := h.UserService.CreateUser(&user)
	if err != nil {
		return responses.DomainError(400, err), err
	}

	return responses.DomainJSON(200, createdUser), nil
}

func (h *UserHandler) updateUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := strings.TrimPrefix(request.Path, "/user/")

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: fmt.Sprintf("Usuário %s atualizado", id)}, nil
}

func (h *UserHandler) getUsers(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var users []*model.User

	usersList, err := h.UserService.ListUsers(users)
	if err != nil {
		return responses.DomainError(400, err), err
	}

	return responses.DomainJSON(200, usersList), nil
}

func (h *UserHandler) getUserById(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := strings.TrimPrefix(request.Path, "/user/")

	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		return responses.DomainError(404, err), err
	}

	return responses.DomainJSON(200, user), nil
}
