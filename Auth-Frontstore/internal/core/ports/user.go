package ports

import (
	domain "github.com/auth-frontstore-service/internal/core/domain"
)

type UserService interface {
	NewUser(username string, password string, role string) (*domain.User, error)
	IsCorrectPassword(user *domain.User, password string) bool
}
