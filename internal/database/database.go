package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// InitializeMongoClient initializes and assigns the global MongoClient
func InitializeMongoClient(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Check connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	MongoClient = client
	return nil
}
