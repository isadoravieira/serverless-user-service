package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/isadoravieira/serverless-user-service/src/internal/domain/model"
)

type UserRepository struct {
	DB *dynamodb.DynamoDB
}

func NewUserRepository(db *dynamodb.DynamoDB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Save(user *model.User) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item: map[string]*dynamodb.AttributeValue{
			"id":        {S: aws.String(user.ID)},
			"name":      {S: aws.String(user.Name)},
			"email":     {S: aws.String(user.Email)},
			"password":  {S: aws.String(user.Password)},
			"createdAt": {S: aws.String(user.CreatedAt)},
			"updatedAt": {S: aws.String(user.UpdatedAt)},
		},
	}

	_, err := r.DB.PutItem(input)
	return err
}

func (r *UserRepository) List() (*dynamodb.ScanOutput, error) {
	items := &dynamodb.ScanInput{
		TableName: aws.String("user"),
	}

	result, err := r.DB.Scan(items)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *UserRepository) GetById(id string) (*model.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("user"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	}

	result, err := r.DB.GetItem(input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, fmt.Errorf("user with ID %s not found", id)
	}

	user := &model.User{
		ID:        *result.Item["id"].S,
		Name:      *result.Item["name"].S,
		Email:     *result.Item["email"].S,
		Password:  *result.Item["password"].S,
		CreatedAt: *result.Item["createdAt"].S,
		UpdatedAt: *result.Item["updatedAt"].S,
	}

	return user, nil
}
