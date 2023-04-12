package initializers

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(collection string) *mongo.Collection {
	return db.Collection(collection)
}

func Connect() error {
	uri := os.Getenv("DB_URL")
	if uri == "" {
		return errors.New("Could not get env variable")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	db = client.Database("SocialApp")

	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}
