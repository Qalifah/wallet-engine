package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func New(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client, nil
}

func CreateTestDB(ctx context.Context) (*mongo.Database, error) {
	testURI := os.Getenv("DATABASE_URI")
	client, err := New(ctx, testURI)
	if err != nil {
		return nil, err
	}
	return client.Database("wallet_engine_test"), nil
}
