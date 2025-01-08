package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DatabaseContract interface {
	CreateUser(ctx context.Context, user User) (*User, error)
	CreateProduct(ctx context.Context, product Product) (*Product, error)
	UpdateProduct(ctx context.Context, productID string, updateParams Product) error
	DeleteProduct(ctx context.Context, productID string) error
}

type MongoRepository struct {
	client      *mongo.Client
	database    *mongo.Database
	userColl    *mongo.Collection
	productColl *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, databaseName string) DatabaseContract {
	db := client.Database(databaseName)
	return &MongoRepository{
		client:      client,
		database:    db,
		userColl:    db.Collection("users"),
		productColl: db.Collection("products"),
	}
}
