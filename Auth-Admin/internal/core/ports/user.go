package ports

import (
	domain "github.com/auth-admin-service/internal/core/domain"
)

type UserRepository interface {
	Get(id string) (*domain.UserPublic, error)
	Save(*domain.User) error
}

type UserService interface {
	NewUser(username string, password string, role string) (*domain.User, error)
	IsCorrectPassword(user *domain.User, password string) bool
}
