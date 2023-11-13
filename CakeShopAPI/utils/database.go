package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO add DB_URI to .env
func ConnectToMongoDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client, nil
}
