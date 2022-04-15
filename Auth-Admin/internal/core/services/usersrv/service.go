package usersrv

import (
	"context"
	"fmt"
	"time"

	domain "github.com/auth-admin-service/internal/core/domain"
	repo "github.com/auth-admin-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
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

func (srv *service) FetchUser(condition *bson.M, user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := repo.ConnectDB().DataStore
	if err := db.Collection("user").FindOne(ctx, condition).Decode(&user); err != nil {
		return err
	}
	return nil
}

func (srv *service) UpdateUser(condition *bson.M, update *bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := repo.ConnectDB().DataStore
	if _, err := db.Collection("user").UpdateOne(ctx, condition, update); err != nil {
		return err
	}
	return nil
}
