package model

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/isadoravieira/serverless-user-service/src/pkg/security"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// prepare will call validate and format methods for user received
func (user *User) PrepareUser(stage string) error {
	if err := user.validateFields(stage); err != nil {
		return err
	}

	if err := user.formatFields(stage); err != nil {
		return err
	}

	return nil
}

func (user *User) validateFields(field string) error {
	if user.Name == "" {
		return errors.New("Name is a required field and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("Email is a required field and cannot be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email entered is an invalid format")
	}

	if field == "register" && user.Password == "" {
		return errors.New("Password is a required field and cannot be blank")
	}

	return nil
}

func (user *User) formatFields(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}
	return nil
}
