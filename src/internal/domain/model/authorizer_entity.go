package model

type Authorizer struct {
	IsEmailVerified bool   `json:"isEmailVerified"`
	IsActive        bool   `json:"isActive"`
	LastLoginAt     string `json:"lastLoginAt"`
	DeactivatedAt   string `json:"deactivatedAt"`
}
