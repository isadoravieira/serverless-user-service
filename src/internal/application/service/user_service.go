package service

import (
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

	dateNow := formatteddates.GetCurrencyFormattedDate()

	user.CreatedAt = dateNow
	user.UpdatedAt = dateNow

	err := s.UserRepo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
