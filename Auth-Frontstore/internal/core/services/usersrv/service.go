package usersrv

import (
	"context"
	"fmt"
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

func (srv *service) VerifyCustomerCredentials(email string, password string) (bool, string) {
	return true, "623cb9f57cbbe99be3193341"
}

func (srv *service) NewUser(role string, customerId primitive.ObjectID) (*domain.User, error) {

	user := &domain.User{
		Role:       role,
		ID:         primitive.NewObjectID(),
		CustomerId: customerId,
		Tokens:     make([]string, 1),
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

func (srv *service) UpdateUser(condition *bson.M, update *bson.M, customerId primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := repo.ConnectDB()
	defer func() {
		db.Client.Disconnect(*db.Context)
	}()
	count, err := db.DataStore.Collection("customers").CountDocuments(ctx, bson.M{"customerId": customerId})
	if err != nil {
		return err
	}
	if count == 0 {
		user, err := srv.NewUser("customer", customerId)
		if err != nil {
			return err
		}
		res, err := db.DataStore.Collection("customers").InsertOne(ctx, user)
		if err != nil {
			return err
		}
		fmt.Println("Inserting New Customer", res)
	}
	if _, err := db.DataStore.Collection("customers").UpdateOne(ctx, condition, update); err != nil {
		return err
	}
	return nil
}
