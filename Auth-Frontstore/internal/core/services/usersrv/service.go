package usersrv

import (
	"context"
	"fmt"
	"time"

	domain "github.com/auth-frontstore-service/internal/core/domain"
	repo "github.com/auth-frontstore-service/internal/repository"
	pb "github.com/auth-frontstore-service/protos/customerpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type service struct{}

func New() *service {
	return &service{}
}

func sendCredentialService(c pb.ServiceClient, username string, password string) (bool, string) {
	credentialRequest := pb.CredentialRequest{
		Credential: &pb.Credential{
			Username: username,
			Password: password,
		},
	}
	fmt.Println(2)
	res, err := c.CredentialService(context.Background(), &credentialRequest)
	if err != nil {
		fmt.Println("Encountered an error while creating the", err.Error())
		return false, err.Error()
	}
	return res.Ispresent, res.Customerid
}

func sendCredential(username string, password string) (bool, string) {
	conn, err := grpc.Dial("35.84.28.237:30215", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println("Error while grpc connection", err.Error())
		return false, err.Error()
	}
	c := pb.NewServiceClient(conn)
	return sendCredentialService(c, username, password)
}

func (srv *service) VerifyCustomerCredentials(email string, password string) (bool, string) {
	return sendCredential(email, password)
}

func (srv *service) NewUser(role string, customerId string) (*domain.User, error) {

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

func (srv *service) UpdateUser(condition *bson.M, update *bson.M, customerId string) error {
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
