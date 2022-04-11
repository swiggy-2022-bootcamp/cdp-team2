package repository

import (
	"context"
	"time"

	"github.com/products-admin-service/internal/core/domain"
	"github.com/products-admin-service/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	mongoClientTimeout = 5
)

type ProductRepository struct {
	MongoClient *mongo.Client
	Database    *mongo.Database
	Collection  *mongo.Collection
}

// Check whether ProductRepository implements IProductRepository
var _ ports.IProductsRepository = (*ProductRepository)(nil)

func NewProductsRepository(conn string) (*ProductRepository, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), mongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		conn,
	))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return &ProductRepository{
		MongoClient: client,
		Database:    client.Database("ProductsAdminDB"),
		Collection:  client.Database("ProductsAdminDB").Collection("products"),
	}, nil
}

func (pr *ProductRepository) AddProduct(_product *domain.Product) error {
	return nil
}

func (pr *ProductRepository) UpdateProduct(_product *domain.Product) error {
	return nil
}

func (pr *ProductRepository) DeleteProduct(_productID primitive.ObjectID) error {
	return nil
}
