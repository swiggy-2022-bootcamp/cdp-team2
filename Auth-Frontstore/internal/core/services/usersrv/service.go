package usersrv

import (
	"context"
	"time"

	domain "github.com/auth-frontstore-service/internal/core/domain"
	repo "github.com/auth-frontstore-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct{}

func New() *service {
	return &service{}
}

func (srv *service) VerifyCustomerCredentials(email string, password string) bool {
	return true
}

func (srv *service) NewUser(role string) (*domain.User, error) {

	user := &domain.User{
		Role: role,
		ID:   primitive.NewObjectID(),
	}

	return user, nil
}

func (srv *service) FetchUser(condition *bson.M, user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := repo.ConnectDB()
	defer func() {
		db.Client.Disconnect(*db.Context)
	}()
	if err := db.DataStore.Collection("customers").FindOne(ctx, condition).Decode(&user); err != nil {
		return err
	}
	return nil
}

func (srv *service) UpdateUser(condition *bson.M, update *bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := repo.ConnectDB()
	defer func() {
		db.Client.Disconnect(*db.Context)
	}()
	if _, err := db.DataStore.Collection("customers").UpdateOne(ctx, condition, update); err != nil {
		return err
	}
	return nil
}
