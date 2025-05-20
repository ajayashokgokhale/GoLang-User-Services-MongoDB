package dbx

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func Mongoinit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	os.Getenv("MONGO_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	mongoClient = client
}

func GetMongoDB() (*mongo.Database, error) {
	if mongoClient == nil {
		return nil, mongo.ErrClientDisconnected
	}
	return mongoClient.Database(os.Getenv("MONGO_DB")), nil
}
