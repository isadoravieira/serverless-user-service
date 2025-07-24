package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"

	"github.com/isadoravieira/serverless-user-service/src/internal/application/repository"
	"github.com/isadoravieira/serverless-user-service/src/internal/domain/model"
	formatteddates "github.com/isadoravieira/serverless-user-service/src/pkg/formatted_dates"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(db *dynamodb.DynamoDB) *UserService {
	return &UserService{
		UserRepo: repository.NewUserRepository(db),
	}
}

func (s *UserService) CreateUser(user *model.User) (*model.User, error) {
	user.ID = uuid.New().String()

	dateNow, err := formatteddates.GetCurrencyFormattedDate()
	if err != nil {
		return nil, err
	}

	user.CreatedAt = dateNow
	user.UpdatedAt = dateNow

	err = s.UserRepo.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) ListUsers(users []*model.User) ([]*model.User, error) {
	resultList, err := s.UserRepo.List()
	if err != nil {
		return nil, err
	}

	for _, item := range resultList.Items {
		users = append(users, &model.User{
			ID:        *item["id"].S,
			Name:      *item["name"].S,
			Email:     *item["email"].S,
			Password:  *item["password"].S,
			CreatedAt: *item["createdAt"].S,
			UpdatedAt: *item["updatedAt"].S,
		})
	}

	return users, nil
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.UserRepo.GetById(id)
}

func (s *UserService) UpdateUser(updateData *model.User) (*model.User, error) {
	existingUser, err := s.UserRepo.GetById(updateData.ID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if updateData.Name != "" {
		existingUser.Name = updateData.Name
	}
	if updateData.Email != "" {
		existingUser.Email = updateData.Email
	}
	if updateData.Password != "" {
		existingUser.Password = updateData.Password
	}

	dateNow, err := formatteddates.GetCurrencyFormattedDate()
	if err != nil {
		return nil, err
	}
	existingUser.UpdatedAt = dateNow

	err = s.UserRepo.Save(existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}
