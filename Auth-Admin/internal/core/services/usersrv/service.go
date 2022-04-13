package usersrv

import (
	"fmt"

	domain "github.com/auth-admin-service/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type service struct{}

func New() *service {
	return &service{}
}

func (srv *service) NewUser(username string, password string, role string) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &domain.User{
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
		ID:             primitive.NewObjectID(),
	}

	return user, nil
}

func (srv *service) IsCorrectPassword(user *domain.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}
